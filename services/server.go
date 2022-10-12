package services

import (
	"log"
	"net"

	"github.com/caiomp87/bots-grpc-api/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBotServiceServer
}

func StartServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("could start listener: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBotServiceServer(s, &Server{})

	log.Println("magic happens on port 50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("could start server: %v", err)
	}
}
