package main

import (
	"context"
	"log"
	"net"

	pb "stateUpdate_project/stateUpdate"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

//estrutura server irá implementar stateUpdate.GreeterServer
type server struct{}

//SayHello implementa stateUpdate.GreeterServer, fazendo com que essa função possa ser do tipo GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil //Retornando
	//referencia para HelloReply, setando valores para o atributo Message
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
