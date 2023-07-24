package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/netip"

	"github.com/gogo/status"
	tclient "go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/apoxy-dev/proximal/core/log"
	serverdb "github.com/apoxy-dev/proximal/server/db"
	sqlc "github.com/apoxy-dev/proximal/server/db/sql"

	endpointv1 "github.com/apoxy-dev/proximal/api/endpoint/v1"
)

type EndpointService struct {
	db *serverdb.DB
	tc tclient.Client
}

func NewEndpointService(db *serverdb.DB, tc tclient.Client) *EndpointService {
	return &EndpointService{
		db: db,
		tc: tc,
	}
}

// isDomainName checks if a string is a presentation-format domain name
// (currently restricted to hostname-compatible "preferred name" LDH labels and
// SRV-like "underscore labels"; see golang.org/issue/12421).
// Taken from https://github.com/golang/go/blob/fe5af1532ab9c749d880c05e0ffe0e17bf874d7f/src/net/dnsclient.go#L75
// Copyright (c) 2009 The Go Authors. All rights reserved.
func isDomainName(s string) bool {
	// The root domain name is valid. See golang.org/issue/45715.
	if s == "." {
		return true
	}

	// See RFC 1035, RFC 3696.
	// Presentation format has dots before every label except the first, and the
	// terminal empty label is optional here because we assume fully-qualified
	// (absolute) input. We must therefore reserve space for the first and last
	// labels' length octets in wire format, where they are necessary and the
	// maximum total length is 255.
	// So our _effective_ maximum is 253, but 254 is not rejected if the last
	// character is a dot.
	l := len(s)
	if l == 0 || l > 254 || l == 254 && s[l-1] != '.' {
		return false
	}

	last := byte('.')
	nonNumeric := false // true once we've seen a letter or hyphen
	partlen := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		default:
			return false
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_':
			nonNumeric = true
			partlen++
		case '0' <= c && c <= '9':
			// fine
			partlen++
		case c == '-':
			// Byte before dash cannot be dot.
			if last == '.' {
				return false
			}
			partlen++
			nonNumeric = true
		case c == '.':
			// Byte before dot cannot be dot, dash.
			if last == '.' || last == '-' {
				return false
			}
			if partlen > 63 || partlen == 0 {
				return false
			}
			partlen = 0
		}
		last = c
	}
	if last == '-' || partlen > 63 {
		return false
	}

	return nonNumeric
}

func (s *EndpointService) validateAddrs(addrs []*endpointv1.Address) (isDomain bool, err error) {
	for i, addr := range addrs {
		isAddrDomain := false
		_, err := netip.ParseAddr(addr.GetHost())
		if err != nil {
			if !isDomainName(addr.GetHost()) {
				return false, fmt.Errorf("not a domain or IP: %v", err)
			}
			isAddrDomain = true
		}
		if isAddrDomain != isDomain {
			if i == 0 {
				isDomain = isAddrDomain
			} else {
				return false, fmt.Errorf("cannot mix domains and IPs")
			}
		}
		if addr.GetPort() == 0 {
			return false, errors.New("port cannot be 0")
		}
	}
	return isDomain, nil
}

func endpointFromRow(row sqlc.Endpoint, defaultUpstream bool, addrs []*endpointv1.Address) *endpointv1.Endpoint {
	return &endpointv1.Endpoint{
		Cluster:         row.Cluster,
		DefaultUpstream: defaultUpstream,
		Addresses:       addrs,
		Status: &endpointv1.EndpointStatus{
			IsDomain: row.IsDomain,
		},
		UseTls:          row.UseTls.Bool,
		DnsLookupFamily: endpointv1.Endpoint_DNSLookupFamily(endpointv1.Endpoint_DNSLookupFamily_value[row.LookupFamily]),
		CreatedAt:       timestamppb.New(row.CreatedAt.Time),
		UpdatedAt:       timestamppb.New(row.UpdatedAt.Time),
	}
}

func dnsLookupFamilyToSQL(family endpointv1.Endpoint_DNSLookupFamily) string {
	switch family {
	case endpointv1.Endpoint_V4_FIRST:
		return "V4_FIRST"
	case endpointv1.Endpoint_V4_ONLY:
		return "V4_ONLY"
	case endpointv1.Endpoint_V6_FIRST:
		return "V6_FIRST"
	case endpointv1.Endpoint_V6_ONLY:
		return "V6_ONLY"
	default:
		return "V4_FIRST"
	}
}

func (s *EndpointService) CreateEndpoint(
	ctx context.Context,
	req *endpointv1.CreateEndpointRequest,
) (*endpointv1.Endpoint, error) {
	log.Infof("CreateEndpoint: %v", req.Endpoint)

	isDomain, err := s.validateAddrs(req.Endpoint.GetAddresses())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if req.Endpoint.GetCluster() == "" {
		return nil, status.Error(codes.InvalidArgument, "cluster cannot be empty")
	}

	_, err = s.db.Queries().GetEndpointByCluster(ctx, req.Endpoint.GetCluster())
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "endpoint already exists")
	}

	tx, err := s.db.Begin()
	if err != nil {
		log.Errorf("failed to begin transaction: %v", err)
		return nil, status.Error(codes.Internal, "failed to create endpoint")
	}
	defer tx.Rollback()
	qtx := s.db.Queries().WithTx(tx)

	e, err := qtx.CreateEndpoint(ctx, sqlc.CreateEndpointParams{
		Cluster:      req.Endpoint.GetCluster(),
		IsDomain:     isDomain,
		UseTls:       sql.NullBool{Bool: req.Endpoint.GetUseTls(), Valid: true},
		LookupFamily: dnsLookupFamilyToSQL(req.Endpoint.GetDnsLookupFamily()),
	})
	if err != nil {
		log.Errorf("failed to create endpoint: %v", err)
		return nil, status.Error(codes.Internal, "failed to create endpoint")
	}

	for _, addr := range req.Endpoint.GetAddresses() {
		_, err := qtx.CreateEndpointAddress(ctx, sqlc.CreateEndpointAddressParams{
			Cluster: e.Cluster,
			Host:    addr.GetHost(),
			Port:    int64(addr.GetPort()),
		})
		if err != nil {
			log.Errorf("failed to create endpoint address: %v", err)
			return nil, status.Error(codes.Internal, "failed to create endpoint")
		}
	}

	n, err := qtx.CountEndpoints(ctx)
	if err != nil {
		log.Errorf("failed to count endpoints: %v", err)
		return nil, status.Error(codes.Internal, "failed to create endpoint")
	}

	// First endpoint is set as default upstream.
	defaultUpstream := n == 1
	if defaultUpstream {
		log.Infof("CreateEndpoint: Create default upstream: %v", e.Cluster)
		if err := qtx.InitDefaultUpstream(ctx, e.Cluster); err != nil {
			log.Errorf("failed to initialize default upstream: %v", err)
			return nil, status.Error(codes.Internal, "failed to create endpoint")
		}
	} else if defaultUpstream = req.Endpoint.GetDefaultUpstream(); defaultUpstream {
		log.Infof("CreateEndpoint: Set default upstream: %v", e.Cluster)
		if err := qtx.SetDefaultUpstream(ctx, e.Cluster); err != nil {
			log.Errorf("failed to set default upstream: %v", err)
			return nil, status.Error(codes.Internal, "failed to create endpoint")
		}
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("failed to commit transaction: %v", err)
		return nil, status.Error(codes.Internal, "failed to create endpoint")
	}

	return endpointFromRow(e, defaultUpstream, req.Endpoint.Addresses), nil
}

func (s *EndpointService) ListEndpoints(
	ctx context.Context,
	req *endpointv1.ListEndpointsRequest,
) (*endpointv1.ListEndpointsResponse, error) {
	resp, err := s.loadEndpoints(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return &endpointv1.ListEndpointsResponse{}, nil
		}
		return nil, status.Error(codes.Internal, "failed to list endpoints")
	}
	return resp, nil
}

func (s *EndpointService) GetEndpoint(ctx context.Context, req *endpointv1.GetEndpointRequest) (*endpointv1.Endpoint, error) {
	e, err := s.db.Queries().GetEndpointByCluster(ctx, req.Cluster)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "endpoint not found")
		}

		log.Errorf("failed to get endpoint: %v", err)
		return nil, status.Error(codes.Internal, "failed to get endpoint")
	}

	addrs, err := s.db.Queries().GetEndpointAddressesByCluster(ctx, req.Cluster)
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Error(codes.Internal, "failed to get endpoint")
	}

	addrpbs := make([]*endpointv1.Address, len(addrs))
	for i, addr := range addrs {
		addrpbs[i] = &endpointv1.Address{
			Host: addr.Host,
			Port: int32(addr.Port),
		}
	}

	du, err := s.db.Queries().GetDefaultUpstream(ctx)
	if err != nil {
		log.Errorf("failed to get default upstream: %v", err)
		return nil, status.Error(codes.Internal, "failed to get endpoint")
	}

	return endpointFromRow(e, e.Cluster == du.Cluster, addrpbs), nil
}

func (s *EndpointService) UpdateEndpoint(
	ctx context.Context,
	req *endpointv1.UpdateEndpointRequest,
) (*endpointv1.Endpoint, error) {
	isDomain, err := s.validateAddrs(req.Endpoint.GetAddresses())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	e, err := s.db.Queries().GetEndpointByCluster(ctx, req.Endpoint.GetCluster())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "endpoint not found")
		}

		log.Errorf("failed to get endpoint: %v", err)
		return nil, status.Error(codes.Internal, "failed to get endpoint")
	}

	tx, err := s.db.Begin()
	if err != nil {
		log.Errorf("failed to begin transaction: %v", err)
		return nil, status.Error(codes.Internal, "failed to update endpoint")
	}
	defer tx.Rollback()
	qtx := s.db.Queries().WithTx(tx)

	addrs, err := qtx.GetEndpointAddressesByCluster(ctx, req.Endpoint.GetCluster())
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Error(codes.Internal, "failed to update endpoint")
	}

	// delete addresses that are not in the request
	for _, addr := range addrs {
		found := false
		for _, a := range req.Endpoint.GetAddresses() {
			if addr.Host == a.GetHost() && addr.Port == int64(a.GetPort()) {
				found = true
				break
			}
		}

		if !found {
			if err := qtx.DeleteEndpointAddress(ctx, sqlc.DeleteEndpointAddressParams{
				Cluster: e.Cluster,
				Host:    addr.Host,
				Port:    addr.Port,
			}); err != nil {
				log.Errorf("failed to delete endpoint address: %v", err)
				return nil, status.Error(codes.Internal, "failed to update endpoint")
			}
		}
	}

	// create addresses that are not in the database
	for _, a := range req.Endpoint.GetAddresses() {
		found := false
		for _, addr := range addrs {
			if addr.Host == a.GetHost() && addr.Port == int64(a.GetPort()) {
				found = true
				break
			}
		}

		if !found {
			if _, err := qtx.CreateEndpointAddress(ctx, sqlc.CreateEndpointAddressParams{
				Cluster: e.Cluster,
				Host:    a.GetHost(),
				Port:    int64(a.GetPort()),
			}); err != nil {
				log.Errorf("failed to create endpoint address: %v", err)
				return nil, status.Error(codes.Internal, "failed to update endpoint")
			}
		}
	}

	if req.Endpoint.DefaultUpstream {
		if err := qtx.SetDefaultUpstream(ctx, e.Cluster); err != nil {
			log.Errorf("failed to set default upstream: %v", err)
			return nil, status.Error(codes.Internal, "failed to update endpoint")
		}
	}

	if e.IsDomain != isDomain {
		if _, err := qtx.UpdateEndpoint(ctx, sqlc.UpdateEndpointParams{
			Cluster:      e.Cluster,
			IsDomain:     isDomain,
			UseTls:       sql.NullBool{Bool: req.Endpoint.GetUseTls(), Valid: true},
			LookupFamily: dnsLookupFamilyToSQL(req.Endpoint.GetDnsLookupFamily()),
		}); err != nil {
			log.Errorf("failed to update endpoint is_domain: %v", err)
			return nil, status.Error(codes.Internal, "failed to update endpoint")
		}
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("failed to commit transaction: %v", err)
		return nil, status.Error(codes.Internal, "failed to update endpoint")
	}

	return endpointFromRow(e, req.Endpoint.DefaultUpstream, req.Endpoint.Addresses), nil
}

func (s *EndpointService) DeleteEndpoint(ctx context.Context, req *endpointv1.DeleteEndpointRequest) (*emptypb.Empty, error) {
	e, err := s.db.Queries().GetEndpointByCluster(ctx, req.Cluster)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "endpoint not found")
		}
	}

	tx, err := s.db.Begin()
	if err != nil {
		log.Errorf("failed to begin transaction: %v", err)
		return nil, status.Error(codes.Internal, "failed to create endpoint")
	}
	defer tx.Rollback()
	qtx := s.db.Queries().WithTx(tx)

	du, err := qtx.GetDefaultUpstream(ctx)
	if err != nil {
		log.Errorf("failed to get default upstream: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete endpoint")
	}
	if e.Cluster == du.Cluster {
		return nil, status.Error(codes.InvalidArgument, "cannot delete default upstream")
	}

	if err := qtx.DeleteEndpoint(ctx, req.Cluster); err != nil {
		log.Errorf("failed to delete endpoint: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete endpoint")
	}

	addrs, err := qtx.GetEndpointAddressesByCluster(ctx, req.Cluster)
	if err != nil && err != sql.ErrNoRows {
		log.Errorf("failed to get endpoint addresses: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete endpoint")
	}

	for _, addr := range addrs {
		if err := qtx.DeleteEndpointAddress(ctx, sqlc.DeleteEndpointAddressParams{
			Cluster: e.Cluster,
			Host:    addr.Host,
			Port:    addr.Port,
		}); err != nil {
			log.Errorf("failed to delete endpoint address: %v", err)
			return nil, status.Error(codes.Internal, "failed to delete endpoint")
		}
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("failed to commit transaction: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete endpoint")
	}

	return &emptypb.Empty{}, nil
}

func (s *EndpointService) loadEndpoints(ctx context.Context) (*endpointv1.ListEndpointsResponse, error) {
	eps, err := s.db.Queries().ListEndpoints(ctx)
	if err != nil {
		log.Errorf("failed to list endpoints: %v", err)
		return nil, err
	}
	du, err := s.db.Queries().GetDefaultUpstream(ctx)
	if err != nil {
		// Don't log error here, it's possible that there is no default upstream
		// if there are no endpoints yet.
		return nil, err
	}
	addrpbs := make(map[string][]*endpointv1.Address)
	addrs, err := s.db.Queries().ListEndpointAddresses(ctx)
	if err != nil {
		log.Errorf("failed to list endpoint addresses: %v", err)
		return nil, err
	}
	for _, a := range addrs {
		addrpbs[a.Cluster] = append(addrpbs[a.Cluster], &endpointv1.Address{
			Host: a.Host,
			Port: int32(a.Port),
		})
	}

	var endpoints []*endpointv1.Endpoint
	for _, e := range eps {
		endpoints = append(endpoints, endpointFromRow(e, e.Cluster == du.Cluster, addrpbs[e.Cluster]))
	}
	return &endpointv1.ListEndpointsResponse{
		Endpoints: endpoints,
	}, nil
}

func (s *EndpointService) InternalListEndpoints(
	ctx context.Context,
	_ *emptypb.Empty,
) (*endpointv1.ListEndpointsResponse, error) {
	resp, err := s.loadEndpoints(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			// Note that other parts of the application assume that at least a
			// default endpoint exists so we must error here.
			return nil, status.Error(codes.NotFound, "no configured endpoints")
		}
		return nil, status.Error(codes.Internal, "failed to list endpoints")
	}
	return resp, nil
}
