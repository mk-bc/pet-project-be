package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func FetchJobCategories(client *JobPortalServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	response, err := client.service.FetchJobCategories(ctx, &pb.FetchJobCategoriesRequest{})
	if err != nil {
		log.Println("Error getting response: ", err)
	}
	log.Println(response)
}
