package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func FetchUserData(client *JobPortalServiceClient, id uint32) *pb.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	in := &pb.FetchUserDataRequest{
		UserId: id,
	}
	response, err := client.service.FetchUserData(ctx, in)
	if err != nil {
		log.Println("Error getting response: ", err)
	}
	log.Println(response)
	return response.User
}
