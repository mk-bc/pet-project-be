package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) FetchCompanyJobsByCategory(
	ctx context.Context,
	req *pb.FetchCompanyJobsByCategoryRequest) (*pb.FetchCompanyJobsByCategoryResponse, error) {
	log.Println("server: fetch jobs of a specific category of a specific company")
	jobs, err := server.Db.FetchCompanyJobsByCategory(req.CompanyId, req.CategoryId)
	if err != nil {
		log.Fatalf("Error retrieving jobs by company of a specific category: %v", err)
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
	return &pb.FetchCompanyJobsByCategoryResponse{
		Jobs: response,
	}, nil
}
