package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func UserCheckSavedJobs(client *JobPortalServiceClient, userID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.UserCheckSavedJobsRequest{
		UserId: userID,
	}

	response, err := client.service.UserCheckSavedJobs(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: user saved jobs: ", err)
	}

	for _, job := range response.Jobs {
		log.Println(job)
	}
}
