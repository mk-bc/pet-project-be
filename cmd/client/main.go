package main

import (
	"log"
	"os"
	"time"

	"github.com/Valgard/godotenv"
	"github.com/mk-bc/pet-project-be/auth"
	"github.com/mk-bc/pet-project-be/client"
	pb "github.com/mk-bc/pet-project-be/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":8080"

const (
	email           = "mk@bcode.in" //"user2@user.in"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {
	const servicePath = "/proto.JobPortalService/"
	return map[string]bool{
		servicePath + "DeleteCompany":              true,
		servicePath + "UpdateCompanyData":          true,
		servicePath + "CreateNewJob":               true,
		servicePath + "CreateNewJobCategory":       true,
		servicePath + "UpdateJobData":              true,
		servicePath + "DeleteJob":                  true,
		servicePath + "FetchApplicantsByJobID":     true,
		servicePath + "ModifyApplicantApplication": true,
		servicePath + "DeleteUser":                 true,
		servicePath + "UpdateUserData":             true,
		servicePath + "UserJobApplication":         true,
		servicePath + "CheckAppliedJobs":           true,
		servicePath + "UserSavedJob":               true,
		servicePath + "UserCheckSavedJobs":         true,
		servicePath + "UserRemoveSavedJob":         true,
	}
}

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error establishing client connection: %v", err)
	}

	err = godotenv.Load("/home/manjunath/Desktop/learning/pet-project/pet-project-be/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	password := os.Getenv("mk_password")
	// password := "user2"

	jobPortalServiceClient := client.NewJobPortalServiceClient(conn, email, password)
	clientInterceptor, err := auth.NewClientInterceptor(jobPortalServiceClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatalf("Error creating interceptor: %v", err)
	}

	interceptedClientCon, err := grpc.Dial(
		"localhost"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(clientInterceptor.Unary()),
	)

	if err != nil {
		log.Fatalf("Error connecting to interceptedClientConn: %v", err)
	}

	grpcClient := client.NewJobPortalServiceClient(interceptedClientCon, email, password)

	// client.ListAllJobs(grpcClient)
	// client.FetchCompanyData(grpcClient, 1) //companyID
	// client.FetchUserData(grpcClient, 1)
	// client.FetchJobCategories(grpcClient)
	// client.FetchJobsByCompanyID(grpcClient, 1)
	// client.FetchJobsByCategoryID(grpcClient, 2)
	// client.FetchCompanyJobsByCategory(grpcClient, 1, 2)
	// client.NewJob(grpcClient, &pb.Job{
	// 	JobTitle:       "Associate Software Engineer",
	// 	JobDescription: "Job description of company",
	// 	Salary:         "50000",
	// 	CompanyId:      1,
	// 	CategoryId:     1,
	// })

	// client.NewCompany(grpcClient,
	// 	&pb.SensitiveData{
	// 		Email:    "ops@google.in",
	// 		Password: "google",
	// 	},
	// 	&pb.Company{
	// 		CompanyName:  "Google",
	// 		CompanyEmail: "ops@google.in",
	// 		CompanyImage: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ-A4V_QmCM2vdqFTxj8e5etw_i31Z27LYGMw&usqp=CAU",
	// 		Description:  "Google description",
	// 	})
	// client.DeleteCompany(grpcClient, 4)
	// updating company data
	// bCode := client.FetchCompanyData(grpcClient, 1)
	// client.UpdateCompanyData(grpcClient,
	// 	&pb.Company{
	// 		CompanyName:  bCode.CompanyName,
	// 		CompanyImage: "https://media.istockphoto.com/id/1139724840/vector/black-column-logo.jpg?s=612x612&w=0&k=20&c=0C98pNvEzyD8P7mSyxj2KrJgfM4G7au7lonOZuqy-sg=",
	// 		Description:  bCode.Description,
	// 	},
	// )
	// creating new job category
	// client.NewJobCategory(grpcClient, &pb.JobCategory{
	// 	CategoryName: "DevOps Engineer",
	// })
	// updating job data
	// client.JobUpdate(grpcClient, &pb.Job{
	// 	JobId:          10,
	// 	JobTitle:       "Assistant to the Manager",
	// 	JobDescription: "Assists the manager but not AssistantManager",
	// 	Salary:         ">50000",
	// })
	// fetching job data
	// client.FetchJob(grpcClient, &pb.Job{
	// 	JobId: 9,
	// })
	// deleting a job posting
	// client.DeleteJob(grpcClient, &pb.DeleteJobRequest{
	// 	JobId: 9,
	// 	Email: email,
	// })
	// fetch applicants by job id
	// client.FetchApplicantsByJobID(grpcClient, 1)
	// modifying applicant application status
	// client.ModifyApplicationStatus(grpcClient, &pb.JobStatus{
	// 	UserId:            1,
	// 	JobId:             1,
	// 	ApplicationStatus: pb.ApplicationStatus_approved,
	// })
	// creating new user
	// client.NewUser(grpcClient,
	// 	&pb.SensitiveData{
	// 		Email:    "user2@user.in",
	// 		Password: "user2",
	// 	},
	// 	&pb.User{
	// 		UserName:    "user2Name",
	// 		UserImage:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ-A4V_QmCM2vdqFTxj8e5etw_i31Z27LYGMw&usqp=CAU",
	// 		Description: "User2 description",
	// 	},
	// )
	// deleting user
	// client.DeleteUser(grpcClient, 5)

	// update user data
	// mk := client.FetchUserData(grpcClient, 1)
	// client.UpdateUserData(grpcClient,
	// 	&pb.User{
	// 		UserName:    mk.UserName,
	// 		UserImage:   mk.UserImage,
	// 		Description: "Update user data rpc call",
	// 	},
	// )

	// user applying for a job
	// client.UserJobApplication(grpcClient, &pb.JobStatus{
	// 	JobId:  2,
	// 	UserId: 1,
	// })

	// user checking up on applied jobs
	// client.CheckUserAppliedJobs(grpcClient, 1)

	// user saving liked jobs
	// client.UserSavedJob(grpcClient, &pb.SavedJob{
	// 	UserId: 6,
	// 	JobId:  9,
	// })

	// user checking liked/saved jobs
	// client.UserCheckSavedJobs(grpcClient, 1)

	// user removing job from liked/saved jobs
	client.UserRemoveSavedJob(grpcClient, &pb.SavedJob{
		UserId: 1,
		JobId:  2,
	})
}
