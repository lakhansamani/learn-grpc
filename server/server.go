package main

import (
	pb "grpc_learning/gen/proto"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("err: ", err)
	}
	server := grpc.NewServer()
	pb.RegisterTestApiServer(server, &testApiServer{})
	err = server.Serve(listener)
	if err != nil {
		log.Println(`server err:`, err)
	}
}
