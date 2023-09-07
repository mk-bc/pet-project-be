package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *JobPortalServiceServer) RegisterUser(
	ctx context.Context,
	req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	log.Println("server: register new user")
	credentials := &models.SensitiveData{
		Email:    req.Credentials.Email,
		Password: req.Credentials.Password,
		Role:     "applicant",
	}

	err := server.Db.RegisterUser(credentials,
		&models.User{
			UserName:    req.User.UserName,
			UserEmail:   req.Credentials.Email,
			UserImage:   req.User.UserImage,
			Description: req.User.Description,
		},
	)
	if err != nil {
		log.Fatal("Error registering user: ", err)
	}
	token, err := server.jwtManager.Generate(credentials)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to generate token: %v", err)
	}
	return &pb.RegisterUserResponse{
		AccessToken: token,
	}, nil
}
