package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	errorPb "github.com/muthu-kumar-u/go-grpc/gen/go/proto/error"
	userPb "github.com/muthu-kumar-u/go-grpc/gen/go/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	serverAddr = "localhost:50051"
	tlsCert    = "certs/tls/cert.pem"     
	privateKey = "certs/jwt/private.pem"
)

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("no PEM block found")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse private key (PKCS#8): %w", err)
	}

	privKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not RSA private key")
	}

	return privKey, nil
}

func GenerateJWT(privateKey *rsa.PrivateKey) (string, error) {
	claims := jwt.MapClaims{
		"sub": "client-123",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"role": "admin",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}


func extractCustomError(err error) {
	st, ok := status.FromError(err)
	if !ok {
		log.Println("Not a gRPC status error")
		return
	}

	for _, detail := range st.Details() {
		if e, ok := detail.(*errorPb.Error); ok {
			log.Printf("Custom error - Code: %s, Message: %s\n", e.Code, e.Message)
		}
	}
	log.Printf("gRPC Error - Code: %s, Message: %s\n", st.Code(), st.Message())
}

func main() {
	creds, err := credentials.NewClientTLSFromFile(tlsCert, "localhost")
	if err != nil {
		log.Fatalf("failed to load TLS cert: %v", err)
	}

	privKey, err := LoadPrivateKey(privateKey)
	if err != nil {
		log.Fatalf("failed to load private key: %v", err)
	}

	token, err := GenerateJWT(privKey)
	if err != nil {
		log.Fatalf("failed to sign JWT: %v", err)
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := userPb.NewUserServiceClient(conn)

	md := metadata.New(map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", token),
	})

	ctxReq, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx := metadata.NewOutgoingContext(ctxReq, md)

	// --- Create ---
	createResp, err := client.Create(ctx, &userPb.CreateRequest{
		Name:   "TestUser", // Try "" here to test custom error
		Email:  "test@demo.com",
		Number: "9876543210",
	})
	if err != nil {
		extractCustomError(err)
		log.Fatalf("Create failed")
	}
	userID := createResp.GetId()
	fmt.Printf("Created User ID: %d\n", userID)

	// --- Read ---
	readResp, err := client.Read(ctx, &userPb.ReadRequest{Id: userID})
	if err != nil {
		extractCustomError(err)
		log.Fatalf("Read failed")
	}
	fmt.Printf("Read User: %+v\n", readResp)

	// --- Update ---
	updateResp, err := client.Update(ctx, &userPb.UpdateRequest{
		Id:     userID,
		Name:   "UpdatedName",
		Email:  "updated@demo.com",
		Number: "1234567890",
	})
	if err != nil {
		extractCustomError(err)
		log.Fatalf("Update failed")
	}
	fmt.Printf("Updated User ID: %d\n", updateResp.GetId())

	// --- Delete ---
	deleteResp, err := client.Delete(ctx, &userPb.DeleteRequest{Id: userID})
	if err != nil {
		extractCustomError(err)
		log.Fatalf("Delete failed")
	}
	fmt.Printf("Deleted User: %s\n", deleteResp.GetMessage())
}