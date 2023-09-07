package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) FetchJobsByCompanyID(
	ctx context.Context,
	req *pb.FetchJobsByCompanyIDRequest) (*pb.FetchJobsByCompanyIDResponse, error) {
	log.Println("server: fetch jobs of a specific company")
	jobs, err := server.Db.FetchJobsByCompanyID(req.CompanyId)
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
	return &pb.FetchJobsByCompanyIDResponse{
		Jobs: response,
	}, nil
}
