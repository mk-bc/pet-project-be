package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func UserJobApplication(client *JobPortalServiceClient, application *pb.JobStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	req := &pb.UserJobApplicationRequest{
		Jobstatus: application,
	}

	response, err := client.service.UserJobApplication(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: user job application: ", err)
	}
	log.Println(response)
}
