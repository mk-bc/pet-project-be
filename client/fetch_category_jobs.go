package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func FetchJobsByCategoryID(client pb.JobPortalServiceClient, id uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	in := &pb.FetchJobsByCategoryIDRequest{
		CategoryId: id,
	}
	response, err := client.FetchJobsByCategoryID(ctx, in)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}
