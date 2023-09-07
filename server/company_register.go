package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *JobPortalServiceServer) RegisterCompany(
	ctx context.Context,
	req *pb.RegisterCompanyRequest) (*pb.RegisterCompanyResponse, error) {
	log.Println("server: registering a new company")
	credentials := &models.SensitiveData{
		Email:    req.CompanyCredentials.Email,
		Password: req.CompanyCredentials.Password,
		Role:     "company",
	}
	err := server.Db.RegisterCompany(credentials,
		&models.Company{
			CompanyName:  req.Company.CompanyName,
			CompanyEmail: req.CompanyCredentials.Email,
			CompanyImage: req.Company.CompanyImage,
			Description:  req.Company.Description,
		})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error registering company: %v", err)
	}

	token, err := server.jwtManager.Generate(credentials)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to generate token: %v", err)
	}
	return &pb.RegisterCompanyResponse{
		AccessToken: token,
	}, nil
}
