package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "stateUpdate_project/stateUpdate"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	//Setando a conexão com o servidor
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn) //esse metodo envia uma instancia do cliente(conn),
	//obtendo como retorno um GreeterClient que possui um atributo que representa
	//um ponteiro para o tipo grpc.ClientConn

	//Fazendo contato com o servidor e imprimindo as respectivas respostas
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
