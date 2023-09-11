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

func TestRegisterCompany(t *testing.T) {
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
		input          *pb.RegisterCompanyRequest
		mockFunc       func()
		expectedOutput *pb.RegisterCompanyResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.RegisterCompanyRequest{
				CompanyCredentials: &pb.SensitiveData{
					Email:    "company@email.com",
					Password: "123345",
				},
				Company: &pb.Company{
					CompanyName:  "Bcode",
					CompanyImage: "image-url",
					Description:  "Company Description",
				},
			},
			mockFunc: func() {
				mockDB.EXPECT().RegisterCompany(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedOutput: &pb.RegisterCompanyResponse{
				AccessToken: "",
			},
			expectedError: nil,
		},
		{
			name: "Invalid credentials",
			input: &pb.RegisterCompanyRequest{
				CompanyCredentials: &pb.SensitiveData{
					Email:    "company@email.com",
					Password: "",
				},
				Company: &pb.Company{
					CompanyName:  "Bcode2",
					CompanyImage: "image-url2",
					Description:  "Company Description2",
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

		_, err := server.RegisterCompany(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		// if !reflect.DeepEqual(got, tc.expectedOutput) {
		// 	t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		// }
	}
}

func TestFetchCompanyData(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.FetchCompanyDataRequest
		mockFunc       func()
		expectedOutput *pb.FetchCompanyDataResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.FetchCompanyDataRequest{
				CompanyId: 1,
			},
			mockFunc: func() {
				mockDB.EXPECT().FetchCompanyData(gomock.Any()).Return(&models.Company{
					CompanyName:  "test-name",
					CompanyEmail: "test@email.com",
					CompanyImage: "image-url",
					Description:  "test description",
				}, nil)
			},
			expectedOutput: &pb.FetchCompanyDataResponse{
				Company: &pb.Company{
					CompanyName:  "test-name",
					CompanyEmail: "test@email.com",
					CompanyImage: "image-url",
					Description:  "test description",
				},
			},
			expectedError: nil,
		},
		{
			name: "Invalid Request",
			input: &pb.FetchCompanyDataRequest{
				CompanyId: 10,
			},
			mockFunc: func() {
				mockDB.EXPECT().FetchCompanyData(gomock.Any()).Return(nil, fmt.Errorf("Record not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Record not found"),
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.FetchCompanyData(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}

func TestFetchApplicantsByJobID(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.FetchApplicantsByJobIDRequest
		mockFunc       func()
		expectedOutput *pb.FetchApplicantsByJobIDResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.FetchApplicantsByJobIDRequest{
				JobId: 1,
			},
			mockFunc: func() {
				mockDB.EXPECT().FetchApplicantsByJobID(gomock.Any()).Return([]*models.JobStatus{
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
			expectedOutput: &pb.FetchApplicantsByJobIDResponse{
				Applicants: []*pb.JobStatus{
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
			expectedError: nil,
		},
		{
			name: "Invalid Request",
			input: &pb.FetchApplicantsByJobIDRequest{
				JobId: 10,
			},
			mockFunc: func() {
				mockDB.EXPECT().FetchApplicantsByJobID(gomock.Any()).Return(nil, fmt.Errorf("Record not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Record not found"),
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.FetchApplicantsByJobID(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}

func TestModifyApplicantApplication(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.ModifyApplicantApplicationRequest
		mockFunc       func()
		expectedOutput *pb.ModifyApplicantApplicationResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.ModifyApplicantApplicationRequest{
				Status: &pb.JobStatus{
					UserId:            1,
					JobId:             1,
					ApplicationStatus: pb.ApplicationStatus_pending,
				},
			},
			mockFunc: func() {
				mockDB.EXPECT().ModifyApplicationStatus(gomock.Any()).Return(&models.JobStatus{
					UserID:            1,
					JobID:             1,
					ApplicationStatus: models.ApplicationStatus("approved"),
				}, nil)
			},
			expectedOutput: &pb.ModifyApplicantApplicationResponse{
				Status: &pb.JobStatus{
					UserId:            1,
					JobId:             1,
					ApplicationStatus: pb.ApplicationStatus_approved,
				},
			},
			expectedError: nil,
		},
		{
			name: "Invalid Request",
			input: &pb.ModifyApplicantApplicationRequest{
				Status: &pb.JobStatus{
					UserId:            1,
					JobId:             100,
					ApplicationStatus: pb.ApplicationStatus_pending,
				},
			},
			mockFunc: func() {
				mockDB.EXPECT().ModifyApplicationStatus(gomock.Any()).Return(nil, fmt.Errorf("Job record not found"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Job record not found"),
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.ModifyApplicantApplication(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}
