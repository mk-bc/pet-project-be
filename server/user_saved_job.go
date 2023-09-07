package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) UserSavedJob(
	ctx context.Context,
	req *pb.UserSavedJobRequest) (*pb.UserSavedJobResponse, error) {
	log.Println("Server: user applying for a job")
	result, err := server.Db.UserSavedJob(models.SavedJob{
		UserID: uint(req.Job.UserId),
		JobID:  uint(req.Job.JobId),
	})
	if err != nil {
		log.Println("Error user applying to a job: ", err)
		return nil, err
	}
	return &pb.UserSavedJobResponse{
		Job: &pb.SavedJob{
			Id:     uint32(result.ID),
			UserId: uint32(result.UserID),
			JobId:  uint32(result.JobID),
		},
	}, nil
}
