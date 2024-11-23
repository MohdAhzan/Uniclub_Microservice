package server

import (
	"log"
	"net"

	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/config"
	"github.com/MohdAhzan/Uniclub_Microservice/INVENTORY_SVC/pkg/pb"
	"google.golang.org/grpc"
)

type ServerGRPC struct {
  server *grpc.Server
  listner net.Listener
}

func NewGrpcServer(cfg config.Config,server pb.InventoryServiceServer ) (*ServerGRPC,error) {
  
  listner, err := net.Listen("tcp", cfg.PORT)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()
  
  pb.RegisterInventoryServiceServer(grpcServer,server)

	return &ServerGRPC{
		server: grpcServer, 
		listner: listner,
	}, nil

}

func (s *ServerGRPC) Start() {
	err := s.server.Serve(s.listner)
	if err != nil {
		log.Fatal("Error starting Grpc Server")
	}
  

}
