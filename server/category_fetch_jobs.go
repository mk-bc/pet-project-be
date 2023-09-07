package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) FetchJobsByCategoryID(
	ctx context.Context,
	req *pb.FetchJobsByCategoryIDRequest) (*pb.FetchJobsByCategoryIDResponse, error) {
	log.Println("server: fetch jobs of a specific category")
	jobs, err := server.Db.FetchJobsByCategoryID(req.CategoryId)
	if err != nil {
		log.Fatalf("Error retrieving jobs by company: %v", err)
	}
	var response []*pb.Job
	for _, job := range jobs {
		response = append(response, &pb.Job{
			JobId:          uint32(job.ID),
			JobTitle:       job.JobTitle,
			JobDescription: job.JobDescription,
			Salary:         job.Salary,
			CompanyId:      uint32(job.CompanyID),
			CategoryId:     uint32(job.CategoryID),
		})
	}
	return &pb.FetchJobsByCategoryIDResponse{
		Jobs: response,
	}, nil
}
