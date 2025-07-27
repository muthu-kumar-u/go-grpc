package interceptor

import (
	"context"
	"crypto/rsa"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Auth(pubKey *rsa.PublicKey) grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req interface{}, nfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
        md, ok := metadata.FromIncomingContext(ctx)
        if !ok || len(md["authorization"]) == 0 {
            return nil, status.Error(codes.Unauthenticated, "missing token")
        }

        tokenStr := md["authorization"][0]

        if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
            tokenStr = tokenStr[7:]
        }

        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return pubKey, nil
        })

        if err != nil || !token.Valid {
            log.Printf("token invalid: %v", err)
            return nil, status.Error(codes.Unauthenticated, "invalid token")
        }

        return handler(ctx, req)
    }
}