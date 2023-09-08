package server

import (
	"context"
	"fmt"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) CreateNewJob(
	ctx context.Context,
	req *pb.CreateNewJobRequest) (*pb.CreateNewJobResponse, error) {
	log.Println("server: create new job")
	job := models.Job{
		JobTitle:       req.Job.JobTitle,
		JobDescription: req.Job.JobDescription,
		Salary:         req.Job.Salary,
		CategoryID:     uint(req.Job.CategoryId),
		CompanyID:      uint(req.Job.CompanyId),
	}
	jobCreated, err := server.Db.CreateNewJob(job)
	if err != nil {
		return nil, fmt.Errorf("Error creating new job: %v", err)
	}
	return &pb.CreateNewJobResponse{
		Job: &pb.Job{
			JobId:          uint32(jobCreated.ID),
			JobTitle:       jobCreated.JobTitle,
			JobDescription: jobCreated.JobDescription,
			Salary:         jobCreated.Salary,
			CategoryId:     uint32(jobCreated.CategoryID),
			CompanyId:      uint32(jobCreated.CompanyID),
		},
	}, nil
}

func (server *JobPortalServiceServer) UpdateJobData(
	ctx context.Context,
	req *pb.UpdateJobDataRequest) (*pb.UpdateJobDataResponse, error) {
	log.Println("server: updating job details")
	err := server.Db.UpdateJobData(models.Job{
		JobTitle:       req.Job.JobTitle,
		JobDescription: req.Job.JobDescription,
		Salary:         req.Job.Salary,
	}, req.Job.JobId)
	if err != nil {
		log.Fatalf("Error updating job details: %v", err)
	}
	return &pb.UpdateJobDataResponse{
		Job: req.Job,
	}, nil
}

func (server *JobPortalServiceServer) DeleteJob(
	ctx context.Context,
	req *pb.DeleteJobRequest) (*pb.DeleteJobResponse, error) {
	log.Println("server: deleting job")
	company, err := server.Db.FetchCompanyDataByEmail(req.Email)
	if err != nil {
		log.Fatalf("Error fetching company data by email")
	}
	job, err := server.Db.FetchJobData(req.JobId)
	if err != nil {
		log.Fatalf("Error fetching job data")
	}
	if job.CompanyID != company.ID {
		log.Fatalf("Unable to delete job posting: credentials mismatch")
	}
	err = server.Db.DeleteJob(req.JobId)
	if err != nil {
		log.Fatalf("Error deleting job: %v", err)
		return &pb.DeleteJobResponse{Status: "Failed"}, err
	}
	return &pb.DeleteJobResponse{Status: "Success"}, nil
}
func (server *JobPortalServiceServer) FetchAllJobs(
	ctx context.Context,
	noParmas *pb.NoParams) (*pb.FetchAllJobsResponse, error) {
	log.Printf("server: fetching all jobs")
	jobs, err := server.Db.FetchAllJobs()
	if err != nil {
		log.Fatalf("Unable to retrieve jobs: %v", err)
	}
	var response []*pb.Job
	for _, job := range jobs {
		response = append(response, &pb.Job{
			JobId:          uint32(job.ID),
			JobTitle:       job.JobTitle,
			JobDescription: job.JobDescription,
			Salary:         job.Salary,
			CompanyId:      uint32(job.CompanyID),
			CategoryId:     uint32(job.CategoryID),
		})
		log.Println(job)
	}

	return &pb.FetchAllJobsResponse{Jobs: response}, nil
}

func (server *JobPortalServiceServer) FetchJobData(
	ctx context.Context,
	req *pb.Job) (*pb.Job, error) {
	log.Printf("server: fetching job data")
	response, err := server.Db.FetchJobData(req.JobId)
	if err != nil {
		log.Fatalf("Error fetching job data from database: %v", err)
		return nil, err
	}
	return &pb.Job{
		JobId:          uint32(response.ID),
		JobTitle:       response.JobTitle,
		JobDescription: response.JobDescription,
		Salary:         response.Salary,
		CompanyId:      uint32(response.CompanyID),
		CategoryId:     uint32(response.CategoryID),
	}, nil
}

func (server *JobPortalServiceServer) FetchJobsByCompanyID(
	ctx context.Context,
	req *pb.FetchJobsByCompanyIDRequest) (*pb.FetchJobsByCompanyIDResponse, error) {
	log.Println("server: fetch jobs of a specific company")
	jobs, err := server.Db.FetchJobsByCompanyID(req.CompanyId)
	if err != nil {
		log.Fatalf("Error retrieving jobs by company: %v", err)
	}
	var response []*pb.Job
	for _, job := range jobs {
		response = append(response, &pb.Job{
			JobId:          uint32(job.ID),
			JobTitle:       job.JobTitle,
			JobDescription: job.JobDescription,
			Salary:         job.Salary,
			CompanyId:      uint32(job.CompanyID),
			CategoryId:     uint32(job.CategoryID),
		})
	}
	return &pb.FetchJobsByCompanyIDResponse{
		Jobs: response,
	}, nil
}

func (server *JobPortalServiceServer) FetchJobsByCategoryID(
	ctx context.Context,
	req *pb.FetchJobsByCategoryIDRequest) (*pb.FetchJobsByCategoryIDResponse, error) {
	log.Println("server: fetch jobs of a specific category")
	jobs, err := server.Db.FetchJobsByCategoryID(req.CategoryId)
	if err != nil {
		log.Fatalf("Error retrieving jobs by company: %v", err)
	}
	var response []*pb.Job
	for _, job := range jobs {
		response = append(response, &pb.Job{
			JobId:          uint32(job.ID),
			JobTitle:       job.JobTitle,
			JobDescription: job.JobDescription,
			Salary:         job.Salary,
			CompanyId:      uint32(job.CompanyID),
			CategoryId:     uint32(job.CategoryID),
		})
	}
	return &pb.FetchJobsByCategoryIDResponse{
		Jobs: response,
	}, nil
}

func (server *JobPortalServiceServer) FetchCompanyJobsByCategory(
	ctx context.Context,
	req *pb.FetchCompanyJobsByCategoryRequest) (*pb.FetchCompanyJobsByCategoryResponse, error) {
	log.Println("server: fetch jobs of a specific category of a specific company")
	jobs, err := server.Db.FetchCompanyJobsByCategory(req.CompanyId, req.CategoryId)
	if err != nil {
		log.Fatalf("Error retrieving jobs by company of a specific category: %v", err)
	}
	var response []*pb.Job
	for _, job := range jobs {
		response = append(response, &pb.Job{
			JobId:          uint32(job.ID),
			JobTitle:       job.JobTitle,
			JobDescription: job.JobDescription,
			Salary:         job.Salary,
			CompanyId:      uint32(job.CompanyID),
			CategoryId:     uint32(job.CategoryID),
		})
	}
	return &pb.FetchCompanyJobsByCategoryResponse{
		Jobs: response,
	}, nil
}
