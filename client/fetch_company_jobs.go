package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func FetchJobsByCompanyID(client pb.JobPortalServiceClient, id uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	in := &pb.FetchJobsByCompanyIDRequest{
		CompanyId: id,
	}
	response, err := client.FetchJobsByCompanyID(ctx, in)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}
