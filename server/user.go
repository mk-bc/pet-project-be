package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *JobPortalServiceServer) RegisterUser(
	ctx context.Context,
	req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	log.Println("server: register new user")
	credentials := &models.SensitiveData{
		Email:    req.Credentials.Email,
		Password: req.Credentials.Password,
		Role:     "applicant",
	}

	err := server.Db.RegisterUser(credentials,
		&models.User{
			UserName:    req.User.UserName,
			UserEmail:   req.Credentials.Email,
			UserImage:   req.User.UserImage,
			Description: req.User.Description,
		},
	)
	if err != nil {
		log.Fatal("Error registering user: ", err)
	}
	token, err := server.jwtManager.Generate(credentials)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to generate token: %v", err)
	}
	return &pb.RegisterUserResponse{
		AccessToken: token,
	}, nil
}

func (server *JobPortalServiceServer) FetchUserData(
	ctx context.Context,
	req *pb.FetchUserDataRequest) (*pb.FetchUserDataResponse, error) {
	log.Println("server: fetch user data")
	user, err := server.Db.FetchUserData(req.UserId)
	if err != nil {
		log.Fatalf("Error retrieving user data: %v", err)
		return nil, err
	}
	return &pb.FetchUserDataResponse{
		User: &pb.User{
			UserId:      uint32(user.ID),
			UserName:    user.UserName,
			UserEmail:   user.UserEmail,
			UserImage:   user.UserImage,
			PhoneNumber: user.PhoneNumber,
			Description: user.Description,
			Skills:      user.Skills,
		},
	}, nil
}

func (server *JobPortalServiceServer) UpdateUserData(
	ctx context.Context,
	req *pb.UpdateUserDataRequest) (*pb.UpdateUserDataResponse, error) {
	log.Println("server: updating user data")
	updatedUser, err := server.Db.UpdateUserData(models.User{
		UserName:    req.User.UserName,
		UserImage:   req.User.UserImage,
		Description: req.User.Description,
	}, models.SensitiveData{
		Email:    req.Credentials.Email,
		Password: req.Credentials.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserDataResponse{
		User: &pb.User{
			UserId:      uint32(updatedUser.ID),
			UserName:    updatedUser.UserName,
			UserEmail:   updatedUser.UserEmail,
			UserImage:   updatedUser.UserImage,
			Description: updatedUser.Description,
		},
	}, nil
}

func (server *JobPortalServiceServer) DeleteUser(
	ctx context.Context,
	req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	log.Println("server: deleting user")
	user, err := server.Db.FetchUserDataByEmail(req.Email)
	if err != nil {
		log.Fatalf("Error getting user data for verification: %v", err)
	}
	if user.ID != uint(req.UserId) {
		log.Fatalf("Unable to delete user: Credentials mismatch")
	}
	err = server.Db.DeleteUser(req.UserId, req.Email)
	if err != nil {
		log.Fatal("Error deleting user data: ", err)
	}
	return &pb.DeleteUserResponse{
		Status: "Deleted user successfully",
	}, nil
}

func (server *JobPortalServiceServer) UserJobApplication(
	ctx context.Context,
	req *pb.UserJobApplicationRequest) (*pb.UserJobApplicationResponse, error) {
	log.Println("Server: user applying for a job")
	result, err := server.Db.UserJobApplication(models.JobStatus{
		UserID:            uint(req.Jobstatus.UserId),
		JobID:             uint(req.Jobstatus.JobId),
		ApplicationStatus: models.ApplicationStatus("pending"),
	})
	if err != nil {
		log.Println("Error user applying to a job: ", err)
		return nil, err
	}
	return &pb.UserJobApplicationResponse{
		Jobstatus: &pb.JobStatus{
			Id:                uint32(result.ID),
			UserId:            uint32(result.UserID),
			JobId:             uint32(result.JobID),
			ApplicationStatus: returnApplicationStatus(result.ApplicationStatus),
		},
	}, nil
}
func (server *JobPortalServiceServer) CheckAppliedJobs(
	ctx context.Context,
	req *pb.CheckAppliedJobsRequest) (*pb.CheckAppliedJobsResponse, error) {
	log.Println("server: user checks applied jobs")
	result, err := server.Db.CheckAppliedJobs(req.UserId)
	if err != nil {
		log.Println("Error getting user applied jobs: ", err)
	}
	var response []*pb.JobStatus
	for _, application := range result {
		response = append(response, &pb.JobStatus{
			Id:                uint32(application.ID),
			UserId:            uint32(application.UserID),
			JobId:             uint32(application.JobID),
			ApplicationStatus: returnApplicationStatus(application.ApplicationStatus),
		})
	}
	return &pb.CheckAppliedJobsResponse{
		Jobs: response,
	}, nil
}

func (server *JobPortalServiceServer) UserSavedJob(
	ctx context.Context,
	req *pb.UserSavedJobRequest) (*pb.UserSavedJobResponse, error) {
	log.Println("Server: user applying for a job")
	result, err := server.Db.UserSavedJob(models.SavedJob{
		UserID: uint(req.Job.UserId),
		JobID:  uint(req.Job.JobId),
	})
	if err != nil {
		log.Println("Error user applying to a job: ", err)
		return nil, err
	}
	return &pb.UserSavedJobResponse{
		Job: &pb.SavedJob{
			Id:     uint32(result.ID),
			UserId: uint32(result.UserID),
			JobId:  uint32(result.JobID),
		},
	}, nil
}

func (server *JobPortalServiceServer) UserCheckSavedJobs(
	ctx context.Context,
	req *pb.UserCheckSavedJobsRequest) (*pb.UserCheckSavedJobsResponse, error) {
	log.Println("server: user checks saved jobs")
	result, err := server.Db.CheckSavedJobs(req.UserId)
	if err != nil {
		log.Println("Error getting user saved jobs: ", err)
	}
	var response []*pb.Job
	for _, job := range result {
		response = append(response, &pb.Job{
			JobId:          uint32(job.ID),
			JobTitle:       job.JobTitle,
			JobDescription: job.JobDescription,
			Salary:         job.Salary,
			CompanyId:      uint32(job.CompanyID),
			CategoryId:     uint32(job.CategoryID),
		})
	}
	return &pb.UserCheckSavedJobsResponse{
		Jobs: response,
	}, nil
}

func (server *JobPortalServiceServer) UserRemoveSavedJob(
	ctx context.Context,
	req *pb.UserRemoveSavedJobRequest) (*pb.UserRemoveSavedJobResponse, error) {
	log.Println("server: user removing a job from saved list")
	err := server.Db.UserRemoveSavedJob(models.SavedJob{
		UserID: uint(req.Job.UserId),
		JobID:  uint(req.Job.JobId),
	})
	if err != nil {
		log.Fatal("Error removing a saved job: ", err)
	}
	return &pb.UserRemoveSavedJobResponse{
		Status: "success",
	}, nil
}
