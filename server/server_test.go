package server

import (
	"context"
	"testing"

	pb "grpc-example/hello"
)

// TestSayHello tests the SayHello RPC method
func TestSayHello(t *testing.T) {
	// Create a new server instance
	s := NewServer()
	
	// Create test context
	ctx := context.Background()
	
	// Test case 1: Basic hello request
	req := &pb.HelloRequest{Name: "Alice"}
	resp, err := s.SayHello(ctx, req)
	
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	
	expectedMsg := "Hello Alice"
	if resp.GetMessage() != expectedMsg {
		t.Errorf("SayHello returned %q, want %q", resp.GetMessage(), expectedMsg)
	}
}

// TestSayGoodbye tests the SayGoodbye RPC method
func TestSayGoodbye(t *testing.T) {
	s := NewServer()
	ctx := context.Background()
	
	req := &pb.HelloRequest{Name: "Bob"}
	resp, err := s.SayGoodbye(ctx, req)
	
	if err != nil {
		t.Fatalf("SayGoodbye failed: %v", err)
	}
	
	expectedMsg := "Goodbye Bob! See you soon!"
	if resp.GetMessage() != expectedMsg {
		t.Errorf("SayGoodbye returned %q, want %q", resp.GetMessage(), expectedMsg)
	}
}

// TestGetServerStatus tests the GetServerStatus RPC method
func TestGetServerStatus(t *testing.T) {
	s := NewServer()
	ctx := context.Background()
	
	req := &pb.StatusRequest{ServiceName: "hello-service"}
	resp, err := s.GetServerStatus(ctx, req)
	
	if err != nil {
		t.Fatalf("GetServerStatus failed: %v", err)
	}
	
	// Verify status field
	if resp.GetStatus() != "running" {
		t.Errorf("Status returned %q, want %q", resp.GetStatus(), "running")
	}
	
	// Verify version field
	if resp.GetVersion() != "1.0.0" {
		t.Errorf("Version returned %q, want %q", resp.GetVersion(), "1.0.0")
	}
	
	// Verify uptime is reasonable (should be close to 0 since we just started)
	uptime := resp.GetUptimeSeconds()
	if uptime < 0 || uptime > 5 { // Allow up to 5 seconds for test execution
		t.Errorf("Uptime returned %d, expected between 0-5 seconds", uptime)
	}
}

// TestMultipleNames tests all methods with different names
func TestMultipleNames(t *testing.T) {
	s := NewServer()
	ctx := context.Background()
	
	tests := []struct {
		name     string
		method   string
		expected string
	}{
		{"Charlie", "SayHello", "Hello Charlie"},
		{"Diana", "SayGoodbye", "Goodbye Diana! See you soon!"},
		{"Eve", "SayHello", "Hello Eve"},
	}
	
	for _, tt := range tests {
		t.Run(tt.method+"_"+tt.name, func(t *testing.T) {
			req := &pb.HelloRequest{Name: tt.name}
			
			var resp *pb.HelloReply
			var err error
			
			switch tt.method {
			case "SayHello":
				resp, err = s.SayHello(ctx, req)
			case "SayGoodbye":
				resp, err = s.SayGoodbye(ctx, req)
			default:
				t.Errorf("Unknown method: %s", tt.method)
				return
			}
			
			if err != nil {
				t.Fatalf("%s failed: %v", tt.method, err)
			}
			
			if resp.GetMessage() != tt.expected {
				t.Errorf("%s returned %q, want %q", tt.method, resp.GetMessage(), tt.expected)
			}
		})
	}
}

// BenchmarkSayHello benchmarks the SayHello method
func BenchmarkSayHello(b *testing.B) {
	s := NewServer()
	ctx := context.Background()
	req := &pb.HelloRequest{Name: "Benchmark"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := s.SayHello(ctx, req)
		if err != nil {
			b.Fatalf("SayHello failed: %v", err)
		}
	}
}
