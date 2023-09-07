package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func CheckUserAppliedJobs(client *JobPortalServiceClient, userID int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.CheckAppliedJobsRequest{
		UserId: uint32(userID),
	}

	response, err := client.service.CheckAppliedJobs(ctx, req)
	if err != nil {
		log.Fatalf("Error getting response: check user applied jobs: %v", err)
	}
	for _, application := range response.Jobs {
		log.Println(application, application.ApplicationStatus)
	}
}
