package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *JobPortalServiceServer) Login(
	ctx context.Context,
	req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println("server: login")
	user, err := server.Db.Login(req.Credentials.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error finding the user: %v", err)
	}
	if user == nil || user.Password != req.Credentials.Password {
		return nil, status.Errorf(codes.NotFound, "Username/password incorrect")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to generate token")
	}
	return &pb.LoginResponse{
		AccessToken: token,
	}, nil
}
