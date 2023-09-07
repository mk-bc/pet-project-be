package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) DeleteJob(
	ctx context.Context,
	req *pb.DeleteJobRequest) (*pb.DeleteJobResponse, error) {
	log.Println("server: deleting job")
	company, err := server.Db.FetchCompanyDataByEmail(req.Email)
	if err != nil {
		log.Fatalf("Error fetching company data by email")
	}
	job, err := server.Db.FetchJobData(req.JobId)
	if err != nil {
		log.Fatalf("Error fetching job data")
	}
	if job.CompanyID != company.ID {
		log.Fatalf("Unable to delete job posting: credentials mismatch")
	}
	err = server.Db.DeleteJob(req.JobId)
	if err != nil {
		log.Fatalf("Error deleting job: %v", err)
		return &pb.DeleteJobResponse{Status: "Failed"}, err
	}
	return &pb.DeleteJobResponse{Status: "Success"}, nil
}
