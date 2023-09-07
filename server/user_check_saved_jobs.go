package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) UserCheckSavedJobs(
	ctx context.Context,
	req *pb.UserCheckSavedJobsRequest) (*pb.UserCheckSavedJobsResponse, error) {
	log.Println("server: user checks saved jobs")
	result, err := server.Db.CheckSavedJobs(req.UserId)
	if err != nil {
		log.Println("Error getting user saved jobs: ", err)
	}
	var response []*pb.Job
	for _, job := range result {
		response = append(response, &pb.Job{
			JobId:          uint32(job.ID),
			JobTitle:       job.JobTitle,
			JobDescription: job.JobDescription,
			Salary:         job.Salary,
			CompanyId:      uint32(job.CompanyID),
			CategoryId:     uint32(job.CategoryID),
		})
	}
	return &pb.UserCheckSavedJobsResponse{
		Jobs: response,
	}, nil
}
