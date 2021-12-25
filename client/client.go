package main

import (
	"context"
	"log"

	pb "grpc_learning/gen/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println("client err:", err)
	}

	client := pb.NewTestApiClient(conn)
	echoRes, err := client.Echo(context.Background(), &pb.ResponseRequest{
		Msg: "Hello world",
	})
	if err != nil {
		log.Println("echo error:", err)
	}

	log.Println("echo res:", echoRes, echoRes.Msg)
}
