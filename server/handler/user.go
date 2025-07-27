package handler

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	errorPb "github.com/muthu-kumar-u/go-grpc/gen/go/proto/error"
	userPb "github.com/muthu-kumar-u/go-grpc/gen/go/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	ID int32
	Name string
	Email string
	Number string
}

type UserServer struct {
	userPb.UnimplementedUserServiceServer
	users map[int32]User
	mu sync.Mutex
	id int32
}

func NewUserServer() *UserServer {
	return &UserServer{
		users: make(map[int32]User),
	}
}

func (s *UserServer) Create(ctx context.Context, req *userPb.CreateRequest) (*userPb.CreateResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := atomic.AddInt32(&s.id, 1)

	if req.GetName() == "" {
		// Create custom error
		customErr := &errorPb.Error{
			Code:    "400",
			Message: "Name is required",
		}

		// Convert it to gRPC error
		st := status.New(codes.InvalidArgument, customErr.Message)
		details, err := st.WithDetails(customErr)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to add error details: %v", err)
		}
		return nil, details.Err()
	}

	user := User{
		ID:   id,
		Name: req.GetName(),
		Email: req.GetEmail(),
		Number: req.GetNumber(),
	}

	s.users[id] = user

	// normal success path
	resp := &userPb.CreateResponse{Id: id}
	return resp, nil
}

func (s *UserServer) Read(ctx context.Context, req *userPb.ReadRequest) (*userPb.ReadResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.Id]
	if !exists {
		customErr := &errorPb.Error{Code: "404", Message: "User not found"}

		st := status.New(codes.NotFound, customErr.Message)
		details, err := st.WithDetails(customErr)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to add error details: %v", err)
		}
		return nil, details.Err()
	}

	return &userPb.ReadResponse{
		Name: user.Name,
		Email: user.Email,
		Number: user.Number,
	}, nil
}

func (s *UserServer) Update(ctx context.Context, req *userPb.UpdateRequest) (*userPb.UpdateResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.Id]
	if !exists {
		customErr := &errorPb.Error{
			Code:    "404",
			Message: fmt.Sprintf("user with id %d not found", req.Id),
		}

		st := status.New(codes.NotFound, customErr.Message)
		details, err := st.WithDetails(customErr)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to attach error details: %v", err)
		}
		return nil, details.Err()
	}

	user.Name = req.GetName()
	user.Email = req.GetEmail()
	user.Number = req.GetNumber()

	s.users[req.Id] = user

	return &userPb.UpdateResponse{Id: req.Id}, nil
}

func (s *UserServer) Delete(ctx context.Context, req *userPb.DeleteRequest) (*userPb.DeleteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[req.Id]; !exists {
		customErr := &errorPb.Error{
			Code:    "404",
			Message: fmt.Sprintf("user with id %d not found", req.Id),
		}

		st := status.New(codes.NotFound, customErr.Message)
		details, err := st.WithDetails(customErr)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to attach error details: %v", err)
		}
		return nil, details.Err()
	}

	delete(s.users, req.Id)

	return &userPb.DeleteResponse{Message: "User deleted successfully"}, nil
}
