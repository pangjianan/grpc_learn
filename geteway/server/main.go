package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"grpc_learn/geteway/proto/echo"
	"log"
	"net"
	"net/http"
)

type server struct {
	echo.UnimplementedEchoServer
}

func (s *server) Say(ctx context.Context, req *echo.SayRequest) (*echo.SayReply, error) {
	return &echo.SayReply{
		Message: "gateway" + req.Name,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatalf("listen fail err=%s", err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &server{})
	go func() {
		log.Fatalln(s.Serve(listen))
	}()

	conn, err := grpc.DialContext(context.Background(), ":8989", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.DialContext fail err=%s", err)
	}
	echoMux := runtime.NewServeMux()
	err = echo.RegisterEchoHandler(context.Background(), echoMux, conn)
	if err != nil {
		log.Fatalf("RegisterEchoHandler fail err=%s", err)
	}
	echoServer := http.Server{Addr: ":8990", Handler: echoMux}
	err = echoServer.ListenAndServe()
	if err != nil {
		log.Fatalf("ListenAndServe fail err=%s", err)
	}
}
