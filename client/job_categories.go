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

func FetchJobCategories(client *JobPortalServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	response, err := client.service.FetchJobCategories(ctx, &pb.FetchJobCategoriesRequest{})
	if err != nil {
		log.Println("Error getting response: ", err)
	}
	log.Println(response)
}
