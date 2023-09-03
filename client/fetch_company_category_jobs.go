package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func FetchCompanyJobsByCategory(client pb.JobPortalServiceClient, companyID uint32, categoryID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	in := &pb.FetchCompanyJobsByCategoryRequest{
		CompanyId:  companyID,
		CategoryId: categoryID,
	}
	response, err := client.FetchCompanyJobsByCategory(ctx, in)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}
