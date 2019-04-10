package main

import (
	"log"
	"net"

	pb "stateUpdate_project/stateUpdate"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

//estrutura server irá implementar stateUpdate.GreeterServer
type Service struct {
	Server *grpc.Server
	Store  Store
	Config *Config
}

func NewService(config *Config) (*Service, error) {
	service := &Service{
		Server: grpc.NewServer(),
		Store:  &StoreImpl{DB: buildDB(config)},
		Config: config,
	}
	return service, nil
}

func main() {
	config := loadConfig()

	service, err := NewService(&config)
	if err != nil {
		glog.Fatalf("failed to launch server: %v", err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterGreeterServer(service.Server, service)
	if err := service.Server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Print("kkk")

}
