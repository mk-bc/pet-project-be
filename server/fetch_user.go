package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) FetchUserData(
	ctx context.Context,
	req *pb.FetchUserDataRequest) (*pb.FetchUserDataResponse, error) {
	log.Println("server: fetch user data")
	user, err := server.Db.FetchUserData(req.UserId)
	if err != nil {
		log.Fatalf("Error retrieving user data: %v", err)
		return nil, err
	}
	return &pb.FetchUserDataResponse{
		User: &pb.User{
			UserId:      uint32(user.ID),
			UserName:    user.UserName,
			UserEmail:   user.UserEmail,
			UserImage:   user.UserImage,
			PhoneNumber: user.PhoneNumber,
			Description: user.Description,
			Skills:      user.Skills,
		},
	}, nil
}
