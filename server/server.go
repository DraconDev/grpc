package server

import (
	"context"
	"log"
	"net"
	"time"

	pb "grpc-example/hello"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
	startTime time.Time
}

func NewServer() *server {
	return &server{
		startTime: time.Now(),
	}
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayGoodbye(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Goodbye request from: %v", in.GetName())
	return &pb.HelloReply{Message: "Goodbye " + in.GetName() + "! See you soon!"}, nil
}

func (s *server) GetServerStatus(ctx context.Context, in *pb.StatusRequest) (*pb.StatusReply, error) {
	log.Printf("Status request for service: %v", in.GetServiceName())
	uptime := time.Since(s.startTime).Seconds()
	return &pb.StatusReply{
		Status:         "running",
		UptimeSeconds:  int64(uptime),
		Version:        "1.0.0",
	}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, NewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
