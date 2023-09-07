package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) DeleteUser(
	ctx context.Context,
	req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	log.Println("server: deleting user")
	user, err := server.Db.FetchUserDataByEmail(req.Email)
	if err != nil {
		log.Fatalf("Error getting user data for verification: %v", err)
	}
	if user.ID != uint(req.UserId) {
		log.Fatalf("Unable to delete user: Credentials mismatch")
	}
	err = server.Db.DeleteUser(req.UserId, req.Email)
	if err != nil {
		log.Fatal("Error deleting user data: ", err)
	}
	return &pb.DeleteUserResponse{
		Status: "Deleted user successfully",
	}, nil
}
