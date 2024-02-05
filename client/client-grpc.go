package client

import (
	"context"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	backoffLinear = 100 * time.Millisecond
)

func NewGRPCClientServiceConn(ctx context.Context, target string) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(backoffLinear)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
	}

	clientGRPCConn, err := grpc.DialContext(
		ctx,
		target,
		// grpc.WithUnaryInterceptor(traceutils.OpenTracingClientInterceptor(manager.GetTracer())),
		// grpc.WithUnaryInterceptor(manager.GetInterceptor()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		return nil, err
	}

	return clientGRPCConn, nil
}
