package interceptor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func Logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		log.Printf("→ %s | %v | %v", info.FullMethod, time.Since(start), err)
		return resp, err
	}
}
