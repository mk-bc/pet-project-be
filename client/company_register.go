package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func NewCompany(client *JobPortalServiceClient, credentials *pb.SensitiveData, company *pb.Company) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.RegisterCompanyRequest{
		CompanyCredentials: credentials,
		Company:            company,
	}

	_, err := client.service.RegisterCompany(ctx, req)
	if err != nil {
		log.Fatalf("Error registering company: %v", err)
	}
	client.email = credentials.Email
	client.password = credentials.Password
	log.Println("new company created successfully")
}
