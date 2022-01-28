package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_learn/helloword/helloword"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial fail err=%s", err)
	}
	defer conn.Close()

	client := helloword.NewHelloWordClient(conn)
	reply, err := client.Say(context.Background(), &helloword.HelloRequest{Name: "ngr"})
	if err != nil {
		log.Fatalf("client.Say fail err=%s", err)
	}
	log.Printf("reply = %+v", reply.GetMessage())
}
