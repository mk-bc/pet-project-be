package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) FetchCompanyData(
	ctx context.Context,
	req *pb.FetchCompanyDataRequest) (*pb.FetchCompanyDataResponse, error) {
	log.Println("server: fetch company data")
	company, err := server.Db.FetchCompanyData(req.CompanyId)
	if err != nil {
		log.Fatalf("Unable to retrieve company data: %v", err)
	}
	return &pb.FetchCompanyDataResponse{
		Company: &pb.Company{
			CompanyId:    uint32(company.ID),
			CompanyName:  company.CompanyName,
			CompanyEmail: company.CompanyEmail,
			CompanyImage: company.CompanyImage,
			Description:  company.Description,
		},
	}, nil
}
