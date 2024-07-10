package sdk

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	backoffLinear = 100 * time.Millisecond
)

type GRPCServer struct {
	Server *grpc.Server
	Port   string
	Host   string
}

func NewGRPCClientConn(target string) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(backoffLinear)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
	}

	clientGRPCConn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		return nil, err
	}

	return clientGRPCConn, nil
}

// Start Start API server
func NewGRPCServer(server *grpc.Server, host, port string) *GRPCServer {
	s := &GRPCServer{}
	s.Host = host
	s.Port = port
	s.Server = server
	return s
}

// Start Start API server
func (s *GRPCServer) Start(wg *sync.WaitGroup) {
	url := fmt.Sprintf(
		"%s:%s",
		s.Host,
		s.Port,
	)
	lis, err := net.Listen("tcp", url)

	fmt.Printf("[ GRPC Service ] started on %s", url)

	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	err = s.Server.Serve(lis)
	if err != nil {
		panic(err)
	}

	wg.Done()
}
