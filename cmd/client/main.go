package main

import (
	"context"
	"log"

	"github.com/LordRadamanthys/grpc_simple_server-client_golang/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Hello Proto grpc",
	}

	res, err := client.RequestMessage(context.Background(), req)

	if err != nil {
		panic(err)
	}

	log.Print(res.GetStatus())
}
