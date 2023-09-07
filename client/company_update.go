package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func UpdateCompanyData(client *JobPortalServiceClient, company *pb.Company) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.UpdateCompanyDataRequest{
		Credentials: &pb.SensitiveData{
			Email:    client.email,
			Password: client.password,
		},
		Company: company,
	}
	response, err := client.service.UpdateCompanyData(ctx, req)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}
