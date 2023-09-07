package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func NewJobCategory(client *JobPortalServiceClient, category *pb.JobCategory) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.CreateNewJobCategoryRequest{
		Category: category,
	}

	response, err := client.service.CreateNewJobCategory(ctx, req)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println("Successfully created new job category", response)
}
