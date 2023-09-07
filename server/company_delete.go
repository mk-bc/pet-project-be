package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) DeleteCompany(
	ctx context.Context,
	req *pb.DeleteCompanyRequest) (*pb.DeleteCompanyResponse, error) {
	log.Println("server: delete company")
	companyData, err := server.Db.FetchCompanyData(req.CompanyId)
	if err != nil {
		return nil, fmt.Errorf("Error fetching company data to delete: %v", err)
	}
	if req.Email != companyData.CompanyEmail {
		return nil, fmt.Errorf("Error deleting company: Credentials not match: %v", err)
	}

	err = server.Db.DeleteCompany(req.CompanyId, companyData.CompanyEmail)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCompanyResponse{
		Status: "success",
	}, nil
}
