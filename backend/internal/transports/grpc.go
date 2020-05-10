package transports

import (
	"backend/api/playground"
	"backend/config"
	"backend/internal/endpoints"
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

func (g *GRPCTransport) Register(pgEndpoints *endpoints.PlaygroundEndpoint) {
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
	fmt.Println("gRPC server started")
	go func() {
		log.Fatal(g.server.Serve(g.lis))
	}()
}
