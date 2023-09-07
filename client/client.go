package client

import (
	pb "github.com/mk-bc/pet-project-be/proto"
	"google.golang.org/grpc"
)

type JobPortalServiceClient struct {
	service  pb.JobPortalServiceClient
	email    string
	password string
}

func NewJobPortalServiceClient(conn *grpc.ClientConn, email, password string) *JobPortalServiceClient {
	client := pb.NewJobPortalServiceClient(conn)
	return &JobPortalServiceClient{client, email, password}
}
