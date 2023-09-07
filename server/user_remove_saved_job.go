package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) UserRemoveSavedJob(
	ctx context.Context,
	req *pb.UserRemoveSavedJobRequest) (*pb.UserRemoveSavedJobResponse, error) {
	log.Println("server: user removing a job from saved list")
	err := server.Db.UserRemoveSavedJob(models.SavedJob{
		UserID: uint(req.Job.UserId),
		JobID:  uint(req.Job.JobId),
	})
	if err != nil {
		log.Fatal("Error removing a saved job: ", err)
	}
	return &pb.UserRemoveSavedJobResponse{
		Status: "success",
	}, nil
}
