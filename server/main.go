package main

import (
	"context"
	"fmt"
	pb "grpc_test/server/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTestServer
}

func (s *server) GetTestResult(ctx context.Context, req *pb.Hello) (*pb.HelloResult, error) {
	return &pb.HelloResult{Message: req.String() + "hello"}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s", "127.0.0.1:9005"))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterTestServer(s, &server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
