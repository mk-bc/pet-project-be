package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) UpdateUserData(
	ctx context.Context,
	req *pb.UpdateUserDataRequest) (*pb.UpdateUserDataResponse, error) {
	log.Println("server: updating user data")
	updatedUser, err := server.Db.UpdateUserData(models.User{
		UserName:    req.User.UserName,
		UserImage:   req.User.UserImage,
		Description: req.User.Description,
	}, models.SensitiveData{
		Email:    req.Credentials.Email,
		Password: req.Credentials.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserDataResponse{
		User: &pb.User{
			UserId:      uint32(updatedUser.ID),
			UserName:    updatedUser.UserName,
			UserEmail:   updatedUser.UserEmail,
			UserImage:   updatedUser.UserImage,
			Description: updatedUser.Description,
		},
	}, nil
}
