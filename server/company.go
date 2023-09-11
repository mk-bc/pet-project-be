package server

import (
	"context"
	"fmt"
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

	if req.CompanyCredentials.Email == "" || req.CompanyCredentials.Password == "" {
		return nil, fmt.Errorf("Enter proper credentials")
	}

	if req.Company.CompanyName == "" || req.Company.Description == "" {
		return nil, fmt.Errorf("Enter the company mandatory details")
	}

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

func (server *JobPortalServiceServer) FetchCompanyData(
	ctx context.Context,
	req *pb.FetchCompanyDataRequest) (*pb.FetchCompanyDataResponse, error) {
	log.Println("server: fetch company data")
	company, err := server.Db.FetchCompanyData(req.CompanyId)
	if err != nil {
		log.Printf("Unable to retrieve company data: %v", err)
		return nil, err
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

func (server *JobPortalServiceServer) FetchApplicantsByJobID(
	ctx context.Context,
	req *pb.FetchApplicantsByJobIDRequest) (*pb.FetchApplicantsByJobIDResponse, error) {
	log.Println("server: fetching applicants by jobID")
	result, err := server.Db.FetchApplicantsByJobID(req.JobId)
	if err != nil {
		log.Print("Error retreiving applicants data: ", err)
		return nil, err
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

func (server *JobPortalServiceServer) ModifyApplicantApplication(
	ctx context.Context,
	req *pb.ModifyApplicantApplicationRequest) (*pb.ModifyApplicantApplicationResponse, error) {
	log.Println("server: modifying applicant status for a job application")
	result, err := server.Db.ModifyApplicationStatus(models.JobStatus{
		JobID:             uint(req.Status.JobId),
		UserID:            uint(req.Status.UserId),
		ApplicationStatus: models.ApplicationStatus(req.Status.ApplicationStatus.String()),
	})
	if err != nil {
		log.Printf("Error modifying job status: %v", err)
		return nil, err
	}
	return &pb.ModifyApplicantApplicationResponse{
		Status: &pb.JobStatus{
			Id:                uint32(result.ID),
			JobId:             uint32(result.JobID),
			UserId:            uint32(result.UserID),
			ApplicationStatus: returnApplicationStatus(result.ApplicationStatus),
		},
	}, nil
}
