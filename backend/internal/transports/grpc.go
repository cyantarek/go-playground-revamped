package transports

import (
	"backend/api/playground"
	"backend/config"
	"backend/internal/endpoints/rpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GRPCTransport struct {
	lis    net.Listener
	server *grpc.Server
}

func (g *GRPCTransport) Server() *grpc.Server {
	return g.server
}

func (g *GRPCTransport) Register(pgEndpoints *rpc.PlaygroundEndpoint) {
	playground.RegisterPlaygroundServer(g.server, pgEndpoints)
}

func BaseGRPCTransport(cfg *config.Config) *GRPCTransport {
	addr := fmt.Sprintf("%s:%s", cfg.Server.Grpc.Host, cfg.Server.Grpc.Port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()

	reflection.Register(gRPCServer)

	return &GRPCTransport{
		lis:    lis,
		server: gRPCServer,
	}
}

func (g *GRPCTransport) Run() {
	log.Println("gRPC server started on", g.lis.Addr())
	go func() {
		log.Fatal(g.server.Serve(g.lis))
	}()
}

func (g *GRPCTransport) Shutdown() {
	log.Println("emergency termination call. terminating gRPC server")
	
	g.server.GracefulStop()
	
	log.Println("gRPC server terminated")
}
