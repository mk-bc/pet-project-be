package client

import (
	"context"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (client *JobPortalServiceClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Credentials: &pb.SensitiveData{
			Email:    client.email,
			Password: client.password,
		},
	}

	response, err := client.service.Login(ctx, req)
	if err != nil {
		return "", nil
	}
	return response.AccessToken, nil
}
