package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	GRPCPort     = flag.Int("grpc_port", 2020, "Port of a gRPC listener.")
	GatewayPort  = flag.Int("grpc_gateway_port", 8080, "Port of a gRPC gateway instance.")
	HTTPPort     = flag.Int("http_port", 8888, "Port of a HTTP listener.")
	ShutdownWait = flag.Duration("shutdown_wait", 15*time.Second, "How long to wait for server connections to drain.")
	CORSAllowAll = flag.Bool("cors_allow_all", true, "Set CORS headers to allow all requests?")
	SlowReplies  = flag.Bool("slow_replies", false, "Make all requests take an extra second.")
	StaticDir    = flag.String("static_dir", "./frontend/build", "Directory to serve static files from.")

	passedHeaders = map[string]struct{}{
		// Add headers here that you need to pass to the gRPC handler. For example:
		// "stripe-signature": {},
	}
)

func slowReplyWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if *SlowReplies {
			time.Sleep(1 * time.Second)
		}
		h.ServeHTTP(w, r)
	})
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func staticFileServer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/v1/") {
			h.ServeHTTP(w, r)
			return
		} else {
			if r.URL.Path == "/" || !fileExists(*StaticDir+"/"+r.URL.Path) {
				http.ServeFile(w, r, *StaticDir+"/index.html")
			} else {
				http.FileServer(http.Dir(*StaticDir)).ServeHTTP(w, r)
			}
		}
	})
}

func corsAllowAllWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" && *CORSAllowAll {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Add("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
			w.WriteHeader(http.StatusNoContent)
			return
		} else if r.Method != "OPTIONS" && *CORSAllowAll {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Credentials", "true")
		}
		h.ServeHTTP(w, r)
	})
}

type GatewayHandlerFn func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

type ApoxyServer struct {
	Mux       *http.ServeMux
	Srv       *http.Server
	Context   context.Context
	ctxCancel context.CancelFunc

	GRPC *grpc.Server

	Gateway *runtime.ServeMux
	GwSrv   *http.Server

	l           sync.Mutex
	terminating bool
}

type serverOptions struct {
	handlers      []GatewayHandlerFn
	passedHeaders map[string]struct{}
}

func defaultServerOptions() *serverOptions {
	return &serverOptions{
		passedHeaders: passedHeaders,
	}
}

// ServerOption sets Apoxy server options.
type ServerOption func(*serverOptions)

// WithHandlers appends handlers to the list of gRPC handlers of the server.
func WithHandlers(handlers ...GatewayHandlerFn) ServerOption {
	return func(opts *serverOptions) {
		opts.handlers = append(opts.handlers, handlers...)
	}
}

// WithPassedHeader enables passing of HTTP header to the corresponding gRPC handler
// via context. The header will be prefixed with grpcserver-.
func WithPassedHeader(key string) ServerOption {
	return func(opts *serverOptions) {
		opts.passedHeaders[strings.ToLower(key)] = struct{}{}
	}
}

// attachPprofHandlers attaches pprof handlers to the server.
func attachPprofHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.HandleFunc("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.HandleFunc("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.HandleFunc("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.HandleFunc("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
}

// NewApoxyServer returns a new gRPC server.
func NewApoxyServer(opts ...ServerOption) *ApoxyServer {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	sOpts := defaultServerOptions()
	for _, o := range opts {
		o(sOpts)
	}

	s := &ApoxyServer{}
	s.Context, s.ctxCancel = context.WithCancel(context.Background())
	s.Mux = http.NewServeMux()
	s.Srv = &http.Server{
		Addr:           fmt.Sprintf(":%d", *HTTPPort),
		Handler:        s.Mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	attachPprofHandlers(s.Mux)

	jsonpb := &runtime.JSONPb{}
	s.Gateway = runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			if _, ok := sOpts.passedHeaders[strings.ToLower(key)]; ok {
				return runtime.MetadataPrefix + key, true
			}
			return runtime.DefaultHeaderMatcher(key)
		}),
	)
	s.GwSrv = &http.Server{
		Addr:           fmt.Sprintf(":%d", *GatewayPort),
		Handler:        staticFileServer(slowReplyWrapper(corsAllowAllWrapper(s.Gateway))),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	gwOpts := []grpc.DialOption{grpc.WithInsecure()}
	for _, h := range sOpts.handlers {
		h(s.Context, s.Gateway, fmt.Sprintf("localhost:%d", *GRPCPort), gwOpts)
	}

	gOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			ServerValidationUnaryInterceptor,
		),
	}
	s.GRPC = grpc.NewServer(gOpts...)

	return s
}

func (s *ApoxyServer) Run() {
	doneCh := make(chan struct{})
	go func() {
		exitCh := make(chan os.Signal, 1) // Buffered because sender is not waiting.
		signal.Notify(exitCh, syscall.SIGTERM)
		select {
		case <-exitCh:
		case <-s.Context.Done():
		}

		sCtx, cancelFn := context.WithTimeout(s.Context, *ShutdownWait)
		defer cancelFn()
		if err := s.GwSrv.Shutdown(sCtx); err != nil {
			// Error from closing listeners, or context timeout.
			log.Printf("gateway server shutdown: %v\n", err)
		}
		if err := s.Srv.Shutdown(sCtx); err != nil {
			log.Printf("HTTP server shutdown: %v\n", err)
		}
		s.GRPC.GracefulStop() // Block until all existing grpc connections return.
		close(doneCh)
	}()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port %d: %v", *GRPCPort, err)
	}
	log.Printf("starting gRPC server on port: %d\n", *GRPCPort)
	go s.GRPC.Serve(l)

	go func() {
		log.Printf("starting http server on port: %d\n", *HTTPPort)
		if err := s.Srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener.
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	log.Printf("starting gateway server on port: %d\n", *GatewayPort)
	if err := s.GwSrv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("gateway server ListenAndServe: %v", err)
	}

	<-doneCh
}

func (s *ApoxyServer) Shutdown() {
	s.l.Lock()
	defer s.l.Unlock()
	if !s.terminating {
		s.ctxCancel()
	}
	s.terminating = true
}
