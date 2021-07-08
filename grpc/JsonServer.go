package grpc

import (
	"bdxcore/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type JsonServer struct {
	port     int
	server   *grpc.Server
	listener net.Listener
}

func (p *JsonServer) Start() {
	var err error
	p.listener, err = net.Listen("tcp", network.GetOutboundAddress(p.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	p.server = grpc.NewServer()
	//pb.RegisterGreeterServer(s, &server{})

	reflection.Register(p.server)
	if err := p.server.Serve(p.listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (p *JsonServer) Stop() {
	if p.server != nil {
		p.server.GracefulStop()
	}

	if p.listener != nil {
		p.listener.Close()
	}
}
