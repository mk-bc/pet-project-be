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

func JobUpdate(client *JobPortalServiceClient, job *pb.Job) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.service.UpdateJobData(ctx, &pb.UpdateJobDataRequest{Job: job})
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}

func DeleteJob(client *JobPortalServiceClient, req *pb.DeleteJobRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.service.DeleteJob(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: ", err)
	}
	log.Println(response)
}

func FetchJob(client *JobPortalServiceClient, job *pb.Job) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.service.FetchJobData(ctx, job)
	if err != nil {
		log.Fatal("Error getting response: ", err)
	}
	log.Println(response)
}

func ListAllJobs(client *JobPortalServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	response, err := client.service.FetchAllJobs(ctx, &pb.NoParams{})
	if err != nil {
		log.Fatalf("Error fetching jobs: %v", err)
	}
	log.Printf("Jobs listed:\n")
	for _, job := range response.Jobs {
		log.Printf("Job: %v", job)
	}
}

func FetchJobsByCompanyID(client *JobPortalServiceClient, id uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	in := &pb.FetchJobsByCompanyIDRequest{
		CompanyId: id,
	}
	response, err := client.service.FetchJobsByCompanyID(ctx, in)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}

func FetchJobsByCategoryID(client *JobPortalServiceClient, id uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	in := &pb.FetchJobsByCategoryIDRequest{
		CategoryId: id,
	}
	response, err := client.service.FetchJobsByCategoryID(ctx, in)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}

func FetchCompanyJobsByCategory(client *JobPortalServiceClient, companyID uint32, categoryID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	in := &pb.FetchCompanyJobsByCategoryRequest{
		CompanyId:  companyID,
		CategoryId: categoryID,
	}
	response, err := client.service.FetchCompanyJobsByCategory(ctx, in)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}
