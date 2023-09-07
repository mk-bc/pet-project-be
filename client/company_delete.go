package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func DeleteCompany(client *JobPortalServiceClient, companyID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	response, err := client.service.DeleteCompany(ctx,
		&pb.DeleteCompanyRequest{CompanyId: companyID, Email: client.email})
	if err != nil {
		log.Fatalf("Unable to delete company record: %v", err)
	}
	log.Println(response.Status)

}
