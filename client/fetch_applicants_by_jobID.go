package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func FetchApplicantsByJobID(client *JobPortalServiceClient, jobID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.service.FetchApplicantsByJobID(ctx, &pb.FetchApplicantsByJobIDRequest{JobId: jobID})
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	for _, i := range response.Applicants {
		log.Println(i)
	}
}
