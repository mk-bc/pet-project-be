package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) FetchApplicantsByJobID(
	ctx context.Context,
	req *pb.FetchApplicantsByJobIDRequest) (*pb.FetchApplicantsByJobIDResponse, error) {
	log.Println("server: fetching applicants by jobID")
	result, err := server.Db.FetchApplicantsByJobID(req.JobId)
	if err != nil {
		log.Fatal("Error retreiving applicants data: ", err)
	}
	var response []*pb.JobStatus
	for _, applicant := range result {
		response = append(response, &pb.JobStatus{
			Id:                uint32(applicant.ID),
			UserId:            uint32(applicant.UserID),
			JobId:             uint32(applicant.JobID),
			ApplicationStatus: returnApplicationStatus(applicant.ApplicationStatus),
		})
	}
	return &pb.FetchApplicantsByJobIDResponse{Applicants: response}, nil
}
