package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) FetchJobData(
	ctx context.Context,
	req *pb.Job) (*pb.Job, error) {
	log.Printf("server: fetching job data")
	response, err := server.Db.FetchJobData(req.JobId)
	if err != nil {
		log.Fatalf("Error fetching job data from database: %v", err)
		return nil, err
	}
	return &pb.Job{
		JobId:          uint32(response.ID),
		JobTitle:       response.JobTitle,
		JobDescription: response.JobDescription,
		Salary:         response.Salary,
		CompanyId:      uint32(response.CompanyID),
		CategoryId:     uint32(response.CategoryID),
	}, nil
}
