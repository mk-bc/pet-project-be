package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func DeleteUser(client *JobPortalServiceClient, userID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	response, err := client.service.DeleteUser(ctx,
		&pb.DeleteUserRequest{UserId: userID, Email: client.email})
	if err != nil {
		log.Fatalf("Unable to delete user record: %v", err)
	}
	log.Println(response.Status)

}
