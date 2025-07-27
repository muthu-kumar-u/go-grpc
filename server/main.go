package main

import (
	"log"
	"net"

	userPb "github.com/muthu-kumar-u/go-grpc/gen/go/proto/user"
	"github.com/muthu-kumar-u/go-grpc/server/config"
	"github.com/muthu-kumar-u/go-grpc/server/global"
	"github.com/muthu-kumar-u/go-grpc/server/handler"
	"github.com/muthu-kumar-u/go-grpc/server/interceptor"
	"github.com/muthu-kumar-u/go-grpc/server/keys"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cfg := config.Load()

	creds, err := credentials.NewServerTLSFromFile(cfg.Server.TLS.Cert, cfg.Server.TLS.Key)
	if err != nil {
		log.Fatalf("failed to load TLS cert/key: %v", err)
	}

	global.PrivatePem, err = keys.LoadPrivateKey(cfg.Auth.JWTPrivate)
	if err != nil {
		log.Fatalf("private key error: %v", err)
	}

	global.PublicPem, err = keys.LoadPublicKey(cfg.Auth.JWTPublic)
	if err != nil {
		log.Fatalf("public key error: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(interceptor.ChainUnary(
			interceptor.Logging(),
			interceptor.Validate(),
			interceptor.Recovery(),
			interceptor.Auth(global.PublicPem), 
		)),
	)

	lis, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userPb.RegisterUserServiceServer(grpcServer, handler.NewUserServer())

	log.Printf("gRPC server running at %s", cfg.Server.Address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}