package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func UpdateUserData(client *JobPortalServiceClient, user *pb.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.UpdateUserDataRequest{
		Credentials: &pb.SensitiveData{
			Email:    client.email,
			Password: client.password,
		},
		User: user,
	}
	response, err := client.service.UpdateUserData(ctx, req)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}
