package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) UpdateJobData(
	ctx context.Context,
	req *pb.UpdateJobDataRequest) (*pb.UpdateJobDataResponse, error) {
	log.Println("server: updating job details")
	err := server.Db.UpdateJobData(models.Job{
		JobTitle:       req.Job.JobTitle,
		JobDescription: req.Job.JobDescription,
		Salary:         req.Job.Salary,
	}, req.Job.JobId)
	if err != nil {
		log.Fatalf("Error updating job details: %v", err)
	}
	return &pb.UpdateJobDataResponse{
		Job: req.Job,
	}, nil
}
