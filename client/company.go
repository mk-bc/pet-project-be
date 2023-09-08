package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func NewCompany(client *JobPortalServiceClient, credentials *pb.SensitiveData, company *pb.Company) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.RegisterCompanyRequest{
		CompanyCredentials: credentials,
		Company:            company,
	}

	_, err := client.service.RegisterCompany(ctx, req)
	if err != nil {
		log.Fatalf("Error registering company: %v", err)
	}
	client.email = credentials.Email
	client.password = credentials.Password
	log.Println("new company created successfully")
}

func UpdateCompanyData(client *JobPortalServiceClient, company *pb.Company) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.UpdateCompanyDataRequest{
		Credentials: &pb.SensitiveData{
			Email:    client.email,
			Password: client.password,
		},
		Company: company,
	}
	response, err := client.service.UpdateCompanyData(ctx, req)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}

func DeleteCompany(client *JobPortalServiceClient, companyID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	response, err := client.service.DeleteCompany(ctx,
		&pb.DeleteCompanyRequest{CompanyId: companyID, Email: client.email})
	if err != nil {
		log.Fatalf("Unable to delete company record: %v", err)
	}
	log.Println(response.Status)

}

func FetchCompanyData(client *JobPortalServiceClient, id int) *pb.Company {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	in := &pb.FetchCompanyDataRequest{
		CompanyId: uint32(id),
	}
	response, err := client.service.FetchCompanyData(ctx, in)
	if err != nil {
		log.Printf("Error getting response: %v\n", err)
	}
	log.Printf("response: %v\n", response)
	return response.Company
}

func FetchApplicantsByJobID(client *JobPortalServiceClient, jobID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.service.FetchApplicantsByJobID(ctx, &pb.FetchApplicantsByJobIDRequest{JobId: jobID})
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	for _, i := range response.Applicants {
		log.Println(i)
	}
}

func ModifyApplicationStatus(client *JobPortalServiceClient, status *pb.JobStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.ModifyApplicantApplicationRequest{
		Status: status,
	}

	response, err := client.service.ModifyApplicantApplication(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: ", err)
	}
	log.Println(response.Status)
	// applicationStatus needs to be read separately
}
