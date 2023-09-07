package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) UserJobApplication(
	ctx context.Context,
	req *pb.UserJobApplicationRequest) (*pb.UserJobApplicationResponse, error) {
	log.Println("Server: user applying for a job")
	result, err := server.Db.UserJobApplication(models.JobStatus{
		UserID:            uint(req.Jobstatus.UserId),
		JobID:             uint(req.Jobstatus.JobId),
		ApplicationStatus: models.ApplicationStatus("pending"),
	})
	if err != nil {
		log.Println("Error user applying to a job: ", err)
		return nil, err
	}
	return &pb.UserJobApplicationResponse{
		Jobstatus: &pb.JobStatus{
			Id:                uint32(result.ID),
			UserId:            uint32(result.UserID),
			JobId:             uint32(result.JobID),
			ApplicationStatus: returnApplicationStatus(result.ApplicationStatus),
		},
	}, nil
}
