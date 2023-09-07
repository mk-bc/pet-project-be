package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func FetchJob(client *JobPortalServiceClient, job *pb.Job) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.service.FetchJobData(ctx, job)
	if err != nil {
		log.Fatal("Error getting response: ", err)
	}
	log.Println(response)
}
