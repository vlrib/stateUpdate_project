package main

import (
	"context"
	"log"
	"time"

	pb "stateUpdate_project/stateUpdate"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// p, err := c.CreateProject(ctx, &pb.ProjectRequest{Name: name, Status: "ativo"})
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// log.Printf("%v", p.Success)
	// log.Printf("projeto: " + p.Message)

	q, err := c.GetProject(ctx, &pb.SearchRequest{Name: name})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("%v", q.Find)
	log.Printf("%v -- %v -- %v", q.Id, q.Name, q.Status)

}
