package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func UserSavedJob(client *JobPortalServiceClient, savedJob *pb.SavedJob) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	req := &pb.UserSavedJobRequest{
		Job: savedJob,
	}

	response, err := client.service.UserSavedJob(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: user job application: ", err)
	}
	log.Println(response)
}
