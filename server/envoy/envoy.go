package envoy

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	accesslogv3 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	accessloggrpcv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/grpc/v3"
	tapv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/common/tap/v3"
	routerv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	httptapv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/tap/v3"
	httpwasmv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/wasm/v3"
	httpproxyv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	tlsv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	wasmv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/wasm/v3"
	clusterservicev3 "github.com/envoyproxy/go-control-plane/envoy/service/cluster/v3"
	discoveryservicev3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	endpointservicev3 "github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	listenerservicev3 "github.com/envoyproxy/go-control-plane/envoy/service/listener/v3"
	routeservicev3 "github.com/envoyproxy/go-control-plane/envoy/service/route/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	xds "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/gogo/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/apoxy-dev/proximal/core/log"

	endpointv1 "github.com/apoxy-dev/proximal/api/endpoint/v1"
	middlewarev1 "github.com/apoxy-dev/proximal/api/middleware/v1"
)

const (
	defaultUpstreamCluster = "default_upstream"
	xdsClusterName         = "xds_cluster"
	alsClusterName         = "als_cluster"
)

// SnapshotManager is responsible for managing the Envoy snapshot cache.
type SnapshotManager struct {
	listenHost   string
	listenPort   int
	syncInterval time.Duration
	buildBaseDir string

	mSvc      middlewarev1.MiddlewareServiceClient
	eSvc      endpointv1.EndpointServiceClient
	xdsServer xds.Server
	cache     cache.SnapshotCache
}

// NewSnapshotManager returns a new *SnapshotManager.
func NewSnapshotManager(
	ctx context.Context,
	mSvc middlewarev1.MiddlewareServiceClient,
	eSvc endpointv1.EndpointServiceClient,
	buildBaseDir string,
	host string, port int,
	syncInterval time.Duration) *SnapshotManager {

	snapshotCache := cache.NewSnapshotCache(false, cache.IDHash{}, nil)
	return &SnapshotManager{
		listenHost:   host,
		listenPort:   port,
		syncInterval: syncInterval,
		mSvc:         mSvc,
		eSvc:         eSvc,
		buildBaseDir: buildBaseDir,
		xdsServer:    xds.NewServer(ctx, snapshotCache, nil),
		cache:        snapshotCache,
	}
}

func dnsLookupFamilyFromProto(f endpointv1.Endpoint_DNSLookupFamily) clusterv3.Cluster_DnsLookupFamily {
	switch f {
	case endpointv1.Endpoint_V4_ONLY:
		return clusterv3.Cluster_V4_ONLY
	case endpointv1.Endpoint_V6_ONLY:
		return clusterv3.Cluster_V6_ONLY
	case endpointv1.Endpoint_V4_FIRST:
		return clusterv3.Cluster_V4_PREFERRED
	case endpointv1.Endpoint_V6_FIRST:
		return clusterv3.Cluster_AUTO
	default:
		return clusterv3.Cluster_AUTO
	}
}

func (s *SnapshotManager) clusterResources(es []*endpointv1.Endpoint) ([]types.Resource, error) {
	var clusters []types.Resource
	var defaultUpstream string
	for _, e := range es {
		log.Debugf("adding cluster: %v", e)

		cl := &clusterv3.Cluster{
			Name:            e.Cluster,
			ConnectTimeout:  durationpb.New(5 * time.Second),
			DnsLookupFamily: dnsLookupFamilyFromProto(e.DnsLookupFamily),
		}
		if e.UseTls {
			for _, a := range e.Addresses {
				tlspb, _ := anypb.New(&tlsv3.UpstreamTlsContext{})
				cl.TransportSocketMatches = append(cl.TransportSocketMatches, &clusterv3.Cluster_TransportSocketMatch{
					Name: a.Host,
					Match: &structpb.Struct{
						Fields: map[string]*structpb.Value{
							"serverName": {
								Kind: &structpb.Value_StringValue{
									StringValue: a.Host,
								},
							},
						},
					},
					TransportSocket: &core.TransportSocket{
						Name: "envoy.transport_sockets.tls",
						ConfigType: &core.TransportSocket_TypedConfig{
							TypedConfig: tlspb,
						},
					},
				})
			}
		}

		if e.Status == nil {
			log.Warnf("endpoint %v has no status", e)
			continue
		}

		if e.Status.IsDomain {
			cl.ClusterDiscoveryType = &clusterv3.Cluster_Type{
				Type: clusterv3.Cluster_STRICT_DNS,
			}
		} else {
			cl.ClusterDiscoveryType = &clusterv3.Cluster_Type{
				Type: clusterv3.Cluster_STATIC,
			}
		}

		cl.LoadAssignment = &endpointv3.ClusterLoadAssignment{
			ClusterName: e.Cluster,
			Endpoints: []*endpointv3.LocalityLbEndpoints{{
				LbEndpoints: make([]*endpointv3.LbEndpoint, len(e.Addresses)),
			}},
		}
		for i, addr := range e.Addresses {
			cl.LoadAssignment.Endpoints[0].LbEndpoints[i] = &endpointv3.LbEndpoint{
				Metadata: &core.Metadata{
					FilterMetadata: map[string]*structpb.Struct{
						"envoy.transport_socket_match": {
							Fields: map[string]*structpb.Value{
								"serverName": {
									Kind: &structpb.Value_StringValue{
										StringValue: addr.Host,
									},
								},
							},
						},
					},
				},
				HostIdentifier: &endpointv3.LbEndpoint_Endpoint{
					Endpoint: &endpointv3.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Address: addr.Host,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: uint32(addr.Port)},
								},
							},
						},
					},
				},
			}
		}

		clusters = append(clusters, cl)

		if e.DefaultUpstream {
			log.Debugf("adding default upstream: %v", e)
			defaultUpstream = e.Cluster
			def := &clusterv3.Cluster{
				Name:                 defaultUpstreamCluster,
				ConnectTimeout:       cl.ConnectTimeout,
				ClusterDiscoveryType: cl.ClusterDiscoveryType,
				DnsLookupFamily:      dnsLookupFamilyFromProto(e.DnsLookupFamily),
				LoadAssignment: &endpointv3.ClusterLoadAssignment{
					ClusterName: defaultUpstreamCluster,
					Endpoints:   cl.LoadAssignment.Endpoints,
				},
				TransportSocketMatches: cl.TransportSocketMatches,
			}
			clusters = append(clusters, def)
		}
	}

	if defaultUpstream == "" {
		return nil, fmt.Errorf("no default upstream found")
	}

	return clusters, nil
}

func (s *SnapshotManager) httpConnectionManager(ctx context.Context, mds []*middlewarev1.Middleware) (*anypb.Any, error) {
	var filters []*httpproxyv3.HttpFilter
	for _, middleware := range mds {
		log.Debugf("adding middleware %s", middleware.Slug)

		switch middleware.Status {
		case middlewarev1.Middleware_READY, middlewarev1.Middleware_PENDING_READY:
		default:
			log.Infof("%s is not ready", middleware.Slug)
			continue
		}

		if middleware.LiveBuildSha == "" {
			log.Warnf("%s has no live build", middleware.Slug)
			continue
		}
		wasmOut := filepath.Join(s.buildBaseDir, middleware.Slug, middleware.LiveBuildSha, "wasm.out")
		if _, err := os.Stat(wasmOut); os.IsNotExist(err) {
			log.Warnf("%s wasm does not exist", middleware.Slug)
			continue
		}

		wasmConfig, err := anypb.New(&wrapperspb.StringValue{
			Value: middleware.RuntimeParams.ConfigString,
		})
		if err != nil {
			return nil, err
		}
		wpb, _ := anypb.New(&httpwasmv3.Wasm{
			Config: &wasmv3.PluginConfig{
				Vm: &wasmv3.PluginConfig_VmConfig{
					VmConfig: &wasmv3.VmConfig{
						Runtime: "envoy.wasm.runtime.v8",
						Code: &core.AsyncDataSource{
							Specifier: &core.AsyncDataSource_Local{
								Local: &core.DataSource{
									Specifier: &core.DataSource_Filename{
										Filename: wasmOut,
									},
								},
							},
						},
					},
				},
				Configuration: wasmConfig,
			},
		})
		filters = append(filters, &httpproxyv3.HttpFilter{
			Name: wellknown.HTTPWasm,
			ConfigType: &httpproxyv3.HttpFilter_TypedConfig{
				TypedConfig: wpb,
			},
		})
	}

	tpb, err := anypb.New(&httptapv3.Tap{
		CommonConfig: &tapv3.CommonExtensionConfig{
			ConfigType: &tapv3.CommonExtensionConfig_AdminConfig{
				AdminConfig: &tapv3.AdminConfig{
					ConfigId: "http_logs",
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	filters = append(filters, &httpproxyv3.HttpFilter{
		Name: resource.APITypePrefix + "envoy.extensions.filters.http.tap.v3.Tap",
		ConfigType: &httpproxyv3.HttpFilter_TypedConfig{
			TypedConfig: tpb,
		},
	})

	rpb, err := anypb.New(&routerv3.Router{})
	if err != nil {
		return nil, err
	}
	filters = append(filters, &httpproxyv3.HttpFilter{
		Name: wellknown.Router,
		ConfigType: &httpproxyv3.HttpFilter_TypedConfig{
			TypedConfig: rpb,
		},
	})

	alspb, err := anypb.New(&accessloggrpcv3.HttpGrpcAccessLogConfig{
		CommonConfig: &accessloggrpcv3.CommonGrpcAccessLogConfig{
			LogName:             "http_log",
			TransportApiVersion: resource.DefaultAPIVersion,
			GrpcService: &core.GrpcService{
				TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &core.GrpcService_EnvoyGrpc{
						ClusterName: alsClusterName,
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	hcm := &httpproxyv3.HttpConnectionManager{
		CodecType:  httpproxyv3.HttpConnectionManager_AUTO,
		StatPrefix: "ingress_http",
		RouteSpecifier: &httpproxyv3.HttpConnectionManager_RouteConfig{
			RouteConfig: &envoy_config_route_v3.RouteConfiguration{
				Name: "local_route",
				VirtualHosts: []*envoy_config_route_v3.VirtualHost{{
					Name:    defaultUpstreamCluster,
					Domains: []string{"*"},
					Routes: []*envoy_config_route_v3.Route{{
						Match: &envoy_config_route_v3.RouteMatch{
							PathSpecifier: &envoy_config_route_v3.RouteMatch_Prefix{
								Prefix: "/",
							},
						},
						Action: &envoy_config_route_v3.Route_Route{
							Route: &envoy_config_route_v3.RouteAction{
								ClusterSpecifier: &envoy_config_route_v3.RouteAction_Cluster{
									Cluster: defaultUpstreamCluster,
								},
								HostRewriteSpecifier: &envoy_config_route_v3.RouteAction_AutoHostRewrite{
									AutoHostRewrite: &wrapperspb.BoolValue{
										Value: true,
									},
								},
							},
						},
					}},
				}},
			},
		},
		HttpFilters: filters,
		AccessLog: []*accesslogv3.AccessLog{{
			Name: wellknown.HTTPGRPCAccessLog,
			ConfigType: &accesslogv3.AccessLog_TypedConfig{
				TypedConfig: alspb,
			},
		}},
	}
	return anypb.New(hcm)
}

func (s *SnapshotManager) listenerResources(ctx context.Context, mds []*middlewarev1.Middleware) ([]types.Resource, error) {
	lst := &listenerv3.Listener{
		Name: "main",
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_TCP,
					Address:  s.listenHost,
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: uint32(s.listenPort),
					},
				},
			},
		},
	}

	hcmpb, err := s.httpConnectionManager(ctx, mds)
	if err != nil {
		return nil, fmt.Errorf("failed to generate http connection manager: %v", err)
	}

	lst.FilterChains = []*listenerv3.FilterChain{{
		Filters: []*listenerv3.Filter{{
			Name: "envoy.filters.network.http_connection_manager",
			ConfigType: &listenerv3.Filter_TypedConfig{
				TypedConfig: hcmpb,
			},
		}},
	}}

	return []types.Resource{lst}, nil
}

func (s *SnapshotManager) sync(ctx context.Context) error {
	id := time.Now().Unix()

	mrsp, err := s.mSvc.InternalList(ctx, &emptypb.Empty{})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Infof("no middlewares found, skipping snapshot (id:%d)", id)
			return nil
		}
		return fmt.Errorf("failed to list middlewares: %v", err)
	}
	ls, err := s.listenerResources(ctx, mrsp.GetMiddlewares())
	if err != nil {
		return err
	}

	ersp, err := s.eSvc.InternalListEndpoints(ctx, &emptypb.Empty{})
	fmt.Println("ersp", ersp)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Infof("no endpoints found, skipping snapshot (id:%d)", id)
			return nil
		}
		return fmt.Errorf("failed to list endpoints: %v", err)
	}
	cls, err := s.clusterResources(ersp.GetEndpoints())
	if err != nil {
		return err
	}

	log.Infof("syncing snapshot (id:%d) for %d endpoints", id, len(ersp.GetEndpoints()))

	snapshot, err := cache.NewSnapshot(
		fmt.Sprintf("%d.0", id),
		map[resource.Type][]types.Resource{
			resource.ClusterType:  cls,
			resource.ListenerType: ls,
		},
	)
	if err != nil {
		return err
	}
	if err := snapshot.Consistent(); err != nil {
		return err
	}

	nodeIDs := s.cache.GetStatusKeys()
	for _, nodeID := range nodeIDs {
		err := s.cache.SetSnapshot(ctx, nodeID, snapshot)
		if err != nil {
			log.Warnf("error setting snapshot for node %s: %v", nodeID, err)
		}
	}

	log.Infof("successfully synced snapshot (id:%d) for %d endpoints", id, len(ersp.GetEndpoints()))

	return nil
}

// Run starts a blocking sync loop that updates the snapshot cache at the
// configured interval.
func (s *SnapshotManager) Run(ctx context.Context) error {
	for {
		select {
		case <-time.After(s.syncInterval):
			if err := s.sync(ctx); err != nil {
				log.Errorf("error syncing snapshot: %v", err)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	panic("unreachable")
}

func (m *SnapshotManager) RegisterXDS(srv *grpc.Server) {
	discoveryservicev3.RegisterAggregatedDiscoveryServiceServer(srv, m.xdsServer)
	endpointservicev3.RegisterEndpointDiscoveryServiceServer(srv, m.xdsServer)
	clusterservicev3.RegisterClusterDiscoveryServiceServer(srv, m.xdsServer)
	routeservicev3.RegisterRouteDiscoveryServiceServer(srv, m.xdsServer)
	listenerservicev3.RegisterListenerDiscoveryServiceServer(srv, m.xdsServer)
}

func (s *SnapshotManager) Shutdown() {
	s.cache = nil
}
