package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func FetchCompanyData(client pb.JobPortalServiceClient, id int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	in := &pb.FetchCompanyDataRequest{
		CompanyId: uint32(id),
	}
	response, err := client.FetchCompanyData(ctx, in)
	if err != nil {
		log.Printf("Error getting response: %v\n", err)
	}
	log.Printf("response: %v\n", response)
}
