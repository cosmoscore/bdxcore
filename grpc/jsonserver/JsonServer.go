package jsonserver

import (
	"context"
	"github.com/cosmoscore/bdxcore/grpc/jsonserver/service"
	"github.com/cosmoscore/bdxcore/network"
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

func NewServer(port int) *JsonServer {
	jsonServer := &JsonServer{}
	jsonServer.port = port

	return jsonServer
}

func (p *JsonServer) Start() {
	var err error
	p.listener, err = net.Listen("tcp", network.GetOutboundAddress(p.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	p.server = grpc.NewServer()
	service.RegisterJsonServiceServer(p.server, &JsonServer{})

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

func (p *JsonServer) Post(context.Context, *service.JsonRequest) (*service.JsonResponse, error) {
	return &service.JsonResponse{Msg: `{"key":"test"}`}, nil
}
