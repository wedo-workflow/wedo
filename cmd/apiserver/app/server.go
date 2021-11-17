package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/wedo-workflow/wedo/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	wedo "github.com/wedo-workflow/wedo/api"
	"github.com/wedo-workflow/wedo/cmd/apiserver/config"
)

type APIServer struct {
	Config  *config.Config
	Runtime *runtime.Runtime
}

func NewAPIServer(config *config.Config) (*APIServer, error) {
	server := &APIServer{
		Config: config,
	}

	rt, err := runtime.NewRuntime(config.RuntimeConfig)
	if err != nil {
		return nil, err
	}
	server.Runtime = rt

	return server, nil
}

func (s *APIServer) Close() {
	fmt.Println("goodbye apiserver!")
}

// Shared options for the logger, with a custom gRPC code to log level function.
var (
	// grpc_recovery.RecoveryHandlerFunc
	grpcRecoveryOpts = []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error(p)
			return fmt.Errorf("panic triggered: %v", p)
		}),
		grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p interface{}) (err error) {
			log.Error(p)
			md, ok := metadata.FromIncomingContext(ctx)
			if ok {
				log.Error(md)
			}
			return fmt.Errorf("context panic triggered: %v", p)
		}),
	}
)

func (s *APIServer) Run() {
	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	logrusEntry := log.NewEntry(log.StandardLogger())
	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)
	serverOptions := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             10 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(grpcRecoveryOpts...),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, []grpc_logrus.Option{
				grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel),
				grpc_logrus.WithCodes(grpc_logging.DefaultErrorToCode),
			}...),
			grpc_logrus.PayloadUnaryServerInterceptor(logrusEntry, func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
				return true
			}),
		),
	}
	grpcServer := grpc.NewServer(serverOptions...)
	defer grpcServer.GracefulStop()

	// Register the server with the runtime.
	wedo.RegisterWedoServiceServer(grpcServer, s)

	go func() {
		lis, err := net.Listen("tcp", s.Config.GRPCEndpoint.String())
		if err != nil {
			log.Fatal("failed to listen: ", err)
		}
		log.Info("Try start to serve grpc at ", s.Config.GRPCEndpoint.String())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("failed to serve: ", err)
		}
	}()

	// Register the gateway with the runtime.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gw := grpc_runtime.NewServeMux(
		grpc_runtime.WithMarshalerOption(
			grpc_runtime.MIMEWildcard,
			&grpc_runtime.JSONPb{
				OrigName:     true,
				EmitDefaults: true,
			},
		),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := wedo.RegisterWedoServiceHandlerFromEndpoint(ctx,
		gw, s.Config.GRPCEndpoint.String(), opts); err != nil {
		log.Fatal("register handler failed: ", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gw)
	mux.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, err := writer.Write([]byte("I'm OK"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})

	httpServer := http.Server{
		Addr:    s.Config.HTTPEndpoint.String(),
		Handler: grpcHandlerFunc(grpcServer, cors(mux)),
	}
	defer func() {
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Warn(err)
		}
	}()
	log.Info("Try http server serve at ", s.Config.HTTPEndpoint.String())
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal("http server serve failed: ", err)
	}
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if allowedOrigin(r.Header.Get("Origin")) {
		// 	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Vary", "Accept-Encoding, Origin")
		// }
		h.ServeHTTP(w, r)
	})
}
