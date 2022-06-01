package main

import (
	"context"
	"log"
	"net"

	"github.com/LordRadamanthys/grpc_simple_server-client_golang/pb"
	"google.golang.org/grpc"
)

func (s *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	log.Print("message receive: ", req.GetMessage())
	response := &pb.Response{
		Status: 1,
	}
	return response, nil
}

type Server struct {
	pb.UnimplementedSendMessageServer
}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSendMessageServer(grpcServer, &Server{})

	port := ":8080"

	listener, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}

}
