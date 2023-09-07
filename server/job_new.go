package server

import (
	"context"
	"fmt"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) CreateNewJob(
	ctx context.Context,
	req *pb.CreateNewJobRequest) (*pb.CreateNewJobResponse, error) {
	log.Println("server: create new job")
	job := models.Job{
		JobTitle:       req.Job.JobTitle,
		JobDescription: req.Job.JobDescription,
		Salary:         req.Job.Salary,
		CategoryID:     uint(req.Job.CategoryId),
		CompanyID:      uint(req.Job.CompanyId),
	}
	jobCreated, err := server.Db.CreateNewJob(job)
	if err != nil {
		return nil, fmt.Errorf("Error creating new job: %v", err)
	}
	return &pb.CreateNewJobResponse{
		Job: &pb.Job{
			JobId:          uint32(jobCreated.ID),
			JobTitle:       jobCreated.JobTitle,
			JobDescription: jobCreated.JobDescription,
			Salary:         jobCreated.Salary,
			CategoryId:     uint32(jobCreated.CategoryID),
			CompanyId:      uint32(jobCreated.CompanyID),
		},
	}, nil
}
