package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func ChainUnary(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	n := len(interceptors)

	if n == 0 {
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			return handler(ctx, req)
		}
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		chain := handler
		for i := n - 1; i >= 0; i-- {
			current := interceptors[i]
			next := chain
			chain = func(c context.Context, r interface{}) (interface{}, error) {
				return current(c, r, info, next)
			}
		}
		return chain(ctx, req)
	}
}
