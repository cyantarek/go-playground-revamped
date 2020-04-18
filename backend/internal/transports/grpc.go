package transports

import (
	"backend/config"
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
