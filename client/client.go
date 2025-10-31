package client

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc-example/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
	defaultName = "world"
)

func StartClient() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Test all available RPC methods
	testAllMethods(c)
}

func testAllMethods(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test 1: SayHello
	name := "Alice"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	
	log.Println("=== Testing SayHello ===")
	r1, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("SayHello response: %s", r1.GetMessage())

	// Test 2: SayGoodbye
	log.Println("=== Testing SayGoodbye ===")
	r2, err := c.SayGoodbye(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("SayGoodbye failed: %v", err)
	}
	log.Printf("SayGoodbye response: %s", r2.GetMessage())

	// Test 3: GetServerStatus
	log.Println("=== Testing GetServerStatus ===")
	r3, err := c.GetServerStatus(ctx, &pb.StatusRequest{ServiceName: "hello-service"})
	if err != nil {
		log.Fatalf("GetServerStatus failed: %v", err)
	}
	log.Printf("Server Status:")
	log.Printf("  Status: %s", r3.GetStatus())
	log.Printf("  Uptime: %d seconds", r3.GetUptimeSeconds())
	log.Printf("  Version: %s", r3.GetVersion())

	log.Println("=== All tests completed successfully! ===")
}
