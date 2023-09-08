package client

import (
	"context"
	"log"
	"time"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func NewUser(client *JobPortalServiceClient, credentials *pb.SensitiveData, user *pb.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.RegisterUserRequest{
		Credentials: credentials,
		User:        user,
	}

	_, err := client.service.RegisterUser(ctx, req)
	if err != nil {
		log.Fatalf("Error registering user: %v", err)
	}
	client.email = credentials.Email
	client.password = credentials.Password
	log.Println("new user created successfully")
}

func UpdateUserData(client *JobPortalServiceClient, user *pb.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.UpdateUserDataRequest{
		Credentials: &pb.SensitiveData{
			Email:    client.email,
			Password: client.password,
		},
		User: user,
	}
	response, err := client.service.UpdateUserData(ctx, req)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
	}
	log.Println(response)
}

func FetchUserData(client *JobPortalServiceClient, id uint32) *pb.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	in := &pb.FetchUserDataRequest{
		UserId: id,
	}
	response, err := client.service.FetchUserData(ctx, in)
	if err != nil {
		log.Println("Error getting response: ", err)
	}
	log.Println(response)
	return response.User
}

func DeleteUser(client *JobPortalServiceClient, userID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	response, err := client.service.DeleteUser(ctx,
		&pb.DeleteUserRequest{UserId: userID, Email: client.email})
	if err != nil {
		log.Fatalf("Unable to delete user record: %v", err)
	}
	log.Println(response.Status)

}

func UserJobApplication(client *JobPortalServiceClient, application *pb.JobStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	req := &pb.UserJobApplicationRequest{
		Jobstatus: application,
	}

	response, err := client.service.UserJobApplication(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: user job application: ", err)
	}
	log.Println(response)
}

func CheckUserAppliedJobs(client *JobPortalServiceClient, userID int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.CheckAppliedJobsRequest{
		UserId: uint32(userID),
	}

	response, err := client.service.CheckAppliedJobs(ctx, req)
	if err != nil {
		log.Fatalf("Error getting response: check user applied jobs: %v", err)
	}
	for _, application := range response.Jobs {
		log.Println(application, application.ApplicationStatus)
	}
}

func UserSavedJob(client *JobPortalServiceClient, savedJob *pb.SavedJob) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	req := &pb.UserSavedJobRequest{
		Job: savedJob,
	}

	response, err := client.service.UserSavedJob(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: user job application: ", err)
	}
	log.Println(response)
}

func UserCheckSavedJobs(client *JobPortalServiceClient, userID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.UserCheckSavedJobsRequest{
		UserId: userID,
	}

	response, err := client.service.UserCheckSavedJobs(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: user saved jobs: ", err)
	}

	for _, job := range response.Jobs {
		log.Println(job)
	}
}

func UserRemoveSavedJob(client *JobPortalServiceClient, savedJob *pb.SavedJob) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	req := &pb.UserRemoveSavedJobRequest{
		Job: savedJob,
	}

	response, err := client.service.UserRemoveSavedJob(ctx, req)
	if err != nil {
		log.Fatal("Error getting response: user removing saved job: ", err)
	}
	log.Println(response)
}
