package server

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/mk-bc/pet-project-be/auth"
	"github.com/mk-bc/pet-project-be/database/data"
	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
	"go.uber.org/mock/gomock"
)

func TestRegisterUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{
		Db:         mockDB,
		jwtManager: auth.NewJWTManager("secret", 5*time.Minute),
	}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.RegisterUserRequest
		mockFunc       func()
		expectedOutput *pb.RegisterUserResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.RegisterUserRequest{
				Credentials: &pb.SensitiveData{
					Email:    "user@email.com",
					Password: "123345",
				},
				User: &pb.User{
					UserName:    "Bcode User",
					UserImage:   "image-url",
					Description: "Company Description",
					PhoneNumber: "999999999",
				},
			},
			mockFunc: func() {
				mockDB.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedOutput: &pb.RegisterUserResponse{
				AccessToken: "",
			},
			expectedError: nil,
		},
		{
			name: "Invalid credentials",
			input: &pb.RegisterUserRequest{
				Credentials: &pb.SensitiveData{
					Email:    "user@email.com",
					Password: "",
				},
				User: &pb.User{
					UserName:    "Bcode user2",
					UserImage:   "image-url2",
					Description: "user Description2",
				},
			},
			mockFunc: func() {
				// no need to run mockFunc since db call isn't going to happen due to invalid data
				// mockDB.EXPECT().RegisterCompany(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Enter proper credentials"),
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		_, err := server.RegisterUser(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		// if !reflect.DeepEqual(got, tc.expectedOutput) {
		// 	t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		// }
	}
}

func TestFetchUserData(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.FetchUserDataRequest
		mockFunc       func()
		expectedOutput *pb.FetchUserDataResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.FetchUserDataRequest{
				UserId: 1,
			},
			mockFunc: func() {
				mockDB.EXPECT().FetchUserData(gomock.Any()).Return(&models.User{
					UserName:    "test-name",
					UserEmail:   "test@email.com",
					UserImage:   "image-url",
					Description: "test description",
					PhoneNumber: "9999999999",
				}, nil)
			},
			expectedOutput: &pb.FetchUserDataResponse{
				User: &pb.User{
					UserName:    "test-name",
					UserEmail:   "test@email.com",
					UserImage:   "image-url",
					Description: "test description",
					PhoneNumber: "9999999999",
				},
			},
			expectedError: nil,
		},
		{
			name: "Invalid Request",
			input: &pb.FetchUserDataRequest{
				UserId: 10,
			},
			mockFunc: func() {
				mockDB.EXPECT().FetchUserData(gomock.Any()).Return(nil, fmt.Errorf("Record not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Record not found"),
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.FetchUserData(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.DeleteUserRequest
		mockFunc       func()
		expectedOutput *pb.DeleteUserResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.DeleteUserRequest{
				UserId: 1,
				Email:  "mk@bcode.in",
			},
			mockFunc: func() {
				mockDB.EXPECT().FetchUserDataByEmail(gomock.Any()).Return(&models.User{}, nil)
				mockDB.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedOutput: &pb.DeleteUserResponse{
				Status: "Deleted User Successfully",
			},
			expectedError: nil,
		},
		{
			name: "Invalid Request",
			input: &pb.DeleteUserRequest{
				UserId: 1,
				Email:  "",
			},
			mockFunc: func() {
				mockDB.EXPECT().FetchUserDataByEmail(gomock.Any()).Return(&models.User{}, nil)
				mockDB.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Return(fmt.Errorf("Record not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide valid email and userID"),
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.DeleteUser(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}

func TestUserJobApplication(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.UserJobApplicationRequest
		mockFunc       func()
		expectedOutput *pb.UserJobApplicationResponse
		expectedError  error
	}{
		{
			name: "Success/Valid testcase",
			input: &pb.UserJobApplicationRequest{
				Jobstatus: &pb.JobStatus{
					JobId:             1,
					UserId:            1,
					ApplicationStatus: pb.ApplicationStatus_pending,
				},
			},
			mockFunc: func() {
				mockDB.EXPECT().UserJobApplication(gomock.Any()).Return(&models.JobStatus{
					JobID:             1,
					UserID:            1,
					ApplicationStatus: models.ApplicationStatus("pending"),
				}, nil)
			},
			expectedOutput: &pb.UserJobApplicationResponse{
				Jobstatus: &pb.JobStatus{
					JobId:             1,
					UserId:            1,
					ApplicationStatus: pb.ApplicationStatus_pending,
				},
			},
			expectedError: nil,
		},
		{
			name: "Invalid testcase",
			input: &pb.UserJobApplicationRequest{
				Jobstatus: &pb.JobStatus{
					JobId:             1,
					UserId:            10,
					ApplicationStatus: pb.ApplicationStatus_pending,
				},
			},
			mockFunc: func() {
				mockDB.EXPECT().UserJobApplication(gomock.Any()).Return(nil, fmt.Errorf("Foreign Key constraint  error"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Foreign Key constraint  error"),
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.UserJobApplication(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}

func TestChecksAppliedJobs(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.CheckAppliedJobsRequest
		mockFunc       func()
		expectedOutput *pb.CheckAppliedJobsResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.CheckAppliedJobsRequest{
				UserId: 1,
			},
			mockFunc: func() {
				mockDB.EXPECT().CheckAppliedJobs(gomock.Any()).Return([]*models.JobStatus{
					{
						JobID:             1,
						UserID:            1,
						ApplicationStatus: models.ApplicationStatus("pending"),
					},
					{
						JobID:             2,
						UserID:            1,
						ApplicationStatus: models.ApplicationStatus("rejected"),
					},
					{
						JobID:             3,
						UserID:            1,
						ApplicationStatus: models.ApplicationStatus("approved"),
					},
				}, nil)
			},
			expectedOutput: &pb.CheckAppliedJobsResponse{
				Jobs: []*pb.JobStatus{
					{
						JobId:             1,
						UserId:            1,
						ApplicationStatus: pb.ApplicationStatus_pending,
					},
					{
						JobId:             2,
						UserId:            1,
						ApplicationStatus: pb.ApplicationStatus_rejected,
					},
					{
						JobId:             3,
						UserId:            1,
						ApplicationStatus: pb.ApplicationStatus_approved,
					},
				},
			},
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.CheckAppliedJobs(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}

func TestUserSavedJob(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.UserSavedJobRequest
		mockFunc       func()
		expectedOutput *pb.UserSavedJobResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.UserSavedJobRequest{
				Job: &pb.SavedJob{
					UserId: 1,
					JobId:  1,
				},
			},
			mockFunc: func() {
				mockDB.EXPECT().UserSavedJob(gomock.Any()).Return(&models.SavedJob{
					UserID: 1,
					JobID:  1,
				}, nil)
			},
			expectedOutput: &pb.UserSavedJobResponse{
				Job: &pb.SavedJob{
					UserId: 1,
					JobId:  1,
				},
			},
			expectedError: nil,
		},
		{
			name: "Invalid",
			input: &pb.UserSavedJobRequest{
				Job: &pb.SavedJob{
					UserId: 100,
					JobId:  1,
				},
			},
			mockFunc: func() {
				mockDB.EXPECT().UserSavedJob(gomock.Any()).Return(nil, fmt.Errorf("Foreign key constraint"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Foreign key constraint"),
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.UserSavedJob(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}

func TestUserCheckSavedJobs(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.UserCheckSavedJobsRequest
		mockFunc       func()
		expectedOutput *pb.UserCheckSavedJobsResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.UserCheckSavedJobsRequest{
				UserId: 1,
			},
			mockFunc: func() {
				mockDB.EXPECT().CheckSavedJobs(gomock.Any()).Return([]*models.Job{
					{
						JobTitle:       "SWE",
						JobDescription: "SWE Desc",
						Salary:         "100000",
						CompanyID:      1,
						CategoryID:     1,
					},
					{
						JobTitle:       "SDET",
						JobDescription: "SDET Desc",
						Salary:         "100000",
						CompanyID:      2,
						CategoryID:     1,
					},
				}, nil)
			},
			expectedOutput: &pb.UserCheckSavedJobsResponse{
				Jobs: []*pb.Job{
					{
						JobTitle:       "SWE",
						JobDescription: "SWE Desc",
						Salary:         "100000",
						CompanyId:      1,
						CategoryId:     1,
					},
					{
						JobTitle:       "SDET",
						JobDescription: "SDET Desc",
						Salary:         "100000",
						CompanyId:      2,
						CategoryId:     1,
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "Invalid",
			input: &pb.UserCheckSavedJobsRequest{
				UserId: 100,
			},
			mockFunc: func() {
				mockDB.EXPECT().CheckSavedJobs(gomock.Any()).Return(nil, fmt.Errorf("Record not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Record not found"),
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.UserCheckSavedJobs(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}
