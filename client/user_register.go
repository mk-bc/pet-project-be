package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func NewUser(client *JobPortalServiceClient, credentials *pb.SensitiveData, user *pb.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.RegisterUserRequest{
		Credentials: credentials,
		User:        user,
	}

	_, err := client.service.RegisterUser(ctx, req)
	if err != nil {
		log.Fatalf("Error registering user: %v", err)
	}
	client.email = credentials.Email
	client.password = credentials.Password
	log.Println("new user created successfully")
}
