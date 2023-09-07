package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func NewJob(client *JobPortalServiceClient, job *pb.Job) {
	log.Println("client: create new job")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	response, err := client.service.CreateNewJob(ctx, &pb.CreateNewJobRequest{Job: job})
	if err != nil {
		log.Fatalf("error creating new job (client): %v\n", err)
	}
	log.Print(response)
}
