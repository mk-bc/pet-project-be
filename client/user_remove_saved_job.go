package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func UserRemoveSavedJob(client *JobPortalServiceClient, savedJob *pb.SavedJob) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	req := &pb.UserRemoveSavedJobRequest{
		Job: savedJob,
	}

	response, err := client.service.UserRemoveSavedJob(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: user removing saved job: ", err)
	}
	log.Println(response)
}
