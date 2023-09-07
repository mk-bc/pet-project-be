package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func ModifyApplicationStatus(client *JobPortalServiceClient, status *pb.JobStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.ModifyApplicantApplicationRequest{
		Status: status,
	}

	response, err := client.service.ModifyApplicantApplication(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: ", err)
	}
	log.Println(response.Status)
	// applicationStatus needs to be read separately
}
