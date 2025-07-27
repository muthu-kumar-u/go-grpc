package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Validate() grpc.UnaryServerInterceptor {
  return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
    if v, ok := req.(interface{Validate() error}); ok {
       if err := v.Validate(); err != nil {
          return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
        }
    }
  
    return handler(ctx, req)
  }
}