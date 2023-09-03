package server

import (
	"github.com/mk-bc/pet-project-be/auth"
	"github.com/mk-bc/pet-project-be/database/data"
	pb "github.com/mk-bc/pet-project-be/proto"
)

type JobPortalServiceServer struct {
	pb.JobPortalServiceServer
	Db         data.Database
	jwtManager *auth.JWTManager
}
