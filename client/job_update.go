package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func JobUpdate(client *JobPortalServiceClient, job *pb.Job) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.service.UpdateJobData(ctx, &pb.UpdateJobDataRequest{Job: job})
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}
