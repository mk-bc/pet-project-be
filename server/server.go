package server

import (
	"github.com/jinzhu/gorm"
	"github.com/mk-bc/pet-project-be/auth"
	"github.com/mk-bc/pet-project-be/database/data"
	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

type JobPortalServiceServer struct {
	pb.JobPortalServiceServer
	Db         data.Database
	jwtManager *auth.JWTManager
}

func NewJobPortalServiceServer(db *gorm.DB, jwtManager *auth.JWTManager) *JobPortalServiceServer {
	return &JobPortalServiceServer{
		Db:         &data.DBClient{Db: db},
		jwtManager: jwtManager,
	}
}

func returnApplicationStatus(status models.ApplicationStatus) pb.ApplicationStatus {
	switch status {
	case "approved":
		return pb.ApplicationStatus_approved
	case "rejected":
		return pb.ApplicationStatus_rejected
	default:
		return pb.ApplicationStatus_pending
	}
}
