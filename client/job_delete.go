package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func DeleteJob(client *JobPortalServiceClient, req *pb.DeleteJobRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.service.DeleteJob(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: ", err)
	}
	log.Println(response)
}
