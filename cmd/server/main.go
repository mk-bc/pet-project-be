package main

import (
	"log"
	"net"
	"time"

	"github.com/mk-bc/pet-project-be/auth"
	"github.com/mk-bc/pet-project-be/database/config"
	pb "github.com/mk-bc/pet-project-be/proto"
	"github.com/mk-bc/pet-project-be/server"
	"google.golang.org/grpc"
)

const port = ":8080"
const secretKey = "secret"
const tokenDuration = 5 * time.Minute

func accessibleRoles() map[string][]string {
	const servicePath = "/proto.JobPortalService/"
	return map[string][]string{
		servicePath + "DeleteCompany":              {"company", "admin"},
		servicePath + "UpdateCompanyData":          {"company"},
		servicePath + "CreateNewJob":               {"company", "admin"},
		servicePath + "CreateNewJobCategory":       {"company", "admin"},
		servicePath + "UpdateJobData":              {"company"},
		servicePath + "DeleteJob":                  {"company", "admin"},
		servicePath + "FetchApplicantsByJobID":     {"company", "admin"},
		servicePath + "ModifyApplicantApplication": {"company"},
		servicePath + "DeleteUser":                 {"applicant", "admin"},
		servicePath + "UpdateUserData":             {"applicant"},
		servicePath + "UserJobApplication":         {"applicant"},
		servicePath + "CheckAppliedJobs":           {"applicant", "admin"},
		servicePath + "UserSavedJob":               {"applicant"},
		servicePath + "UserCheckSavedJobs":         {"applicant", "admin"},
		servicePath + "UserRemoveSavedJob":         {"applicant"},
	}
	// modified user => applicant
}

func main() {

	db := config.DBSetup()
	// port := flag.String("port", "8080", "opening-port")
	// flag.Parse()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	jwt := auth.NewJWTManager(secretKey, tokenDuration)
	interceptor := auth.NewServerInterceptor(jwt, accessibleRoles())

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
	)

	jobPortalServiceServer := server.NewJobPortalServiceServer(db, jwt)
	pb.RegisterJobPortalServiceServer(grpcServer, jobPortalServiceServer)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server on port 8080: %v", err)
	}

	log.Printf("Started server on port 8080")
}
