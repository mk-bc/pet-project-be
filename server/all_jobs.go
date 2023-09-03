package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) FetchAllJobs(
	ctx context.Context,
	noParmas *pb.NoParams) (*pb.FetchAllJobsResponse, error) {
	log.Printf("server: fetching all jobs")
	jobs, err := server.Db.FetchAllJobs()
	if err != nil {
		log.Fatalf("Unable to retrieve jobs: %v", err)
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
		log.Println(job)
	}

	return &pb.FetchAllJobsResponse{Jobs: response}, nil
}
