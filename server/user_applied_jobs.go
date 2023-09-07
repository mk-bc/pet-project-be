package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) CheckAppliedJobs(
	ctx context.Context,
	req *pb.CheckAppliedJobsRequest) (*pb.CheckAppliedJobsResponse, error) {
	log.Println("server: user checks applied jobs")
	result, err := server.Db.CheckAppliedJobs(req.UserId)
	if err != nil {
		log.Println("Error getting user applied jobs: ", err)
	}
	var response []*pb.JobStatus
	for _, application := range result {
		response = append(response, &pb.JobStatus{
			Id:                uint32(application.ID),
			UserId:            uint32(application.UserID),
			JobId:             uint32(application.JobID),
			ApplicationStatus: returnApplicationStatus(application.ApplicationStatus),
		})
	}
	return &pb.CheckAppliedJobsResponse{
		Jobs: response,
	}, nil
}
