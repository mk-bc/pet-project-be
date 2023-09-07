package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func ListAllJobs(client *JobPortalServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	response, err := client.service.FetchAllJobs(ctx, &pb.NoParams{})
	if err != nil {
		log.Fatalf("Error fetching jobs: %v", err)
	}
	log.Printf("Jobs listed:\n")
	for _, job := range response.Jobs {
		log.Printf("Job: %v", job)
	}
}
