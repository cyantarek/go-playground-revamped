package grpc

import (
	"backend/internal/pkg/middlewares"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"backend/config"
	"backend/internal/pkg/logger"
)

type GRPC struct {
	lis    net.Listener
	server *grpc.Server
}

func (g *GRPC) Server() *grpc.Server {
	return g.server
}

func New(mws *middlewares.Middleware) (*GRPC, error) {
	addr := fmt.Sprintf(":%s", config.Cfg.GRPCPort)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	var gRPCServer *grpc.Server

	if config.Cfg.TLSCertFile != "" && config.Cfg.TLSKeyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(config.Cfg.TLSCertFile, config.Cfg.TLSKeyFile)
		if err != nil {
			return nil, err
		}

		gRPCServer = grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(mws.JwtInterceptorGRPC))
	} else {
		gRPCServer = grpc.NewServer(grpc.UnaryInterceptor(mws.JwtInterceptorGRPC))
	}

	reflection.Register(gRPCServer)

	return &GRPC{
		lis:    lis,
		server: gRPCServer,
	}, nil
}

func (g *GRPC) Run() {
	log.Println("gRPC server started on", g.lis.Addr())
	go func() {
		log.Fatal(g.server.Serve(g.lis))
	}()
}

func (g *GRPC) Shutdown(ctx context.Context) {
	logger.Log.Println("emergency termination call. terminating gRPC server")

	g.server.GracefulStop()

	logger.Log.Println("gRPC server terminated")
}
