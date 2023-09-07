package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) ModifyApplicantApplication(
	ctx context.Context,
	req *pb.ModifyApplicantApplicationRequest) (*pb.ModifyApplicantApplicationResponse, error) {
	log.Println("server: modifying applicant status for a job application")
	result, err := server.Db.ModifyApplicationStatus(models.JobStatus{
		JobID:             uint(req.Status.JobId),
		UserID:            uint(req.Status.UserId),
		ApplicationStatus: models.ApplicationStatus(req.Status.ApplicationStatus.String()),
	})
	if err != nil {
		log.Fatalf("Error modifying job status: %v", err)
	}
	return &pb.ModifyApplicantApplicationResponse{
		Status: &pb.JobStatus{
			Id:                uint32(result.ID),
			JobId:             uint32(result.JobID),
			UserId:            uint32(result.UserID),
			ApplicationStatus: returnApplicationStatus(result.ApplicationStatus),
		},
	}, nil
}
