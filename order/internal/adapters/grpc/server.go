package grpc

import (
	"fmt"
	"log"
	"net"

	"Github.com/a-samir97/microservices/order/config"
	"Github.com/a-samir97/microservices/order/internal/ports"
	"github.com/huseyinbabal/microservices-proto/golang/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		log.Fatalf("Failed to listen on port %d", a.port)
	}

	grpcServer := grpc.NewServer()

	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve grpc on port %d", a.port)
	}
}
