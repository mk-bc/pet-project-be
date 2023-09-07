package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) UpdateCompanyData(
	ctx context.Context,
	req *pb.UpdateCompanyDataRequest) (*pb.UpdatecompanyDataResponse, error) {
	log.Println("server: updating company data")
	updatedCompany, err := server.Db.UpdateCompanyData(models.Company{
		CompanyName:  req.Company.CompanyName,
		CompanyImage: req.Company.CompanyImage,
		Description:  req.Company.Description,
	}, models.SensitiveData{
		Email:    req.Credentials.Email,
		Password: req.Credentials.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdatecompanyDataResponse{
		Status: "success",
		Company: &pb.Company{
			CompanyId:    uint32(updatedCompany.ID),
			CompanyName:  updatedCompany.CompanyName,
			CompanyEmail: updatedCompany.CompanyEmail,
			CompanyImage: updatedCompany.CompanyImage,
			Description:  updatedCompany.Description,
		},
	}, nil
}
