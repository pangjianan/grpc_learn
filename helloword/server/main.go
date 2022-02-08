package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_learn/helloword/helloword"
	"log"
	"net"
)

type Server struct {
	helloword.UnimplementedHelloWordServer
}

func (s *Server) Say(ctx context.Context, request *helloword.HelloRequest) (*helloword.HelloReply, error) {
	log.Printf("receive name=%s", request.GetName())
	return &helloword.HelloReply{
		Message: fmt.Sprintf("你好，%s!", request.GetName()),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50002))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloword.RegisterHelloWordServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
