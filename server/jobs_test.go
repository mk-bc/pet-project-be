package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/mk-bc/pet-project-be/database/data"
	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
	"go.uber.org/mock/gomock"
)

// func TestFetchJobData(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()

// 	mockDB := data.NewMockDatabase(controller)
// 	server := &JobPortalServiceServer{Db: mockDB}

// 	ctx := context.Background()

// 	mockDB.EXPECT().FetchJobData(gomock.Any()).Return(&models.Job{
// 		JobTitle:       "Software Engineer",
// 		JobDescription: "Description",
// 		Salary:         "50000",
// 		CompanyID:      1,
// 		CategoryID:     1,
// 	}, nil)

// 	expected := &pb.Job{
// 		JobTitle:       "Software Engineer",
// 		JobDescription: "Description",
// 		Salary:         "50000",
// 		CompanyId:      1,
// 		CategoryId:     1,
// 	}

// 	got, err := server.FetchJobData(ctx, &pb.Job{JobId: 1})
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// 	if !reflect.DeepEqual(expected, got) {
// 		t.Errorf("Error:\nExpected: %v\nGot: %v\n", expected, got)
// 	}

// }

func TestCreateNewJob(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.CreateNewJobRequest
		mockFunc       func()
		expectedOutput *pb.CreateNewJobResponse
		expectedError  error
	}{
		{
			name: "Success/Valid testcase",
			input: &pb.CreateNewJobRequest{
				Job: &pb.Job{
					JobTitle:       "SWE",
					JobDescription: "SWE Description",
					Salary:         "50000",
					CategoryId:     1,
					CompanyId:      1,
				},
			},
			mockFunc: func() {
				mockDB.EXPECT().CreateNewJob(gomock.Any()).Return(&models.Job{
					JobTitle:       "SWE",
					JobDescription: "SWE Description",
					Salary:         "50000",
					CategoryID:     1,
					CompanyID:      1,
				}, nil)
			},
			expectedOutput: &pb.CreateNewJobResponse{
				Job: &pb.Job{
					JobTitle:       "SWE",
					JobDescription: "SWE Description",
					Salary:         "50000",
					CategoryId:     1,
					CompanyId:      1,
				},
				Status: "successfully created new job posting",
			},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.CreateNewJob(ctx, tc.input)

		if err != tc.expectedError {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got.Job, tc.expectedOutput.Job) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}

func TestFetchJobData(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.Job
		mockFunc       func()
		expectedOutput *pb.Job
		expectedError  error
	}{
		{
			name:  "Success/Valid case",
			input: &pb.Job{JobId: 1},
			mockFunc: func() {
				mockDB.EXPECT().FetchJobData(gomock.Any()).Return(&models.Job{
					JobTitle:       "Software Engineer",
					JobDescription: "Description",
					Salary:         "50000",
					CompanyID:      1,
					CategoryID:     1,
				}, nil)
			},
			expectedOutput: &pb.Job{
				JobTitle:       "Software Engineer",
				JobDescription: "Description",
				Salary:         "50000",
				CompanyId:      1,
				CategoryId:     1,
			},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mockFunc()
		got, err := server.FetchJobData(ctx, tc.input)

		if err != tc.expectedError {
			t.Errorf("Unexpected error:\nExpected: %v\nGot: %v\n", tc.expectedError, err)
		}
		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unexpected output:\nExpected: %v\nGot: %v\n", tc.expectedOutput, got)
		}
	}

}

func TestFetchAlllJobs(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		mockFunc       func()
		expectedOutput []*pb.Job
		expectedError  error
	}{
		{
			name: "Success - Fetches all jobs",
			mockFunc: func() {
				mockDB.EXPECT().FetchAllJobs().Return([]*models.Job{
					{
						JobTitle:       "SWE",
						JobDescription: "SWE Desc",
						Salary:         "100000",
						CompanyID:      1,
						CategoryID:     1,
					},
					{
						JobTitle:       "HR",
						JobDescription: "HR Desc",
						Salary:         "100000",
						CompanyID:      1,
						CategoryID:     2,
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
			expectedOutput: []*pb.Job{
				{
					JobTitle:       "SWE",
					JobDescription: "SWE Desc",
					Salary:         "100000",
					CompanyId:      1,
					CategoryId:     1,
				},
				{
					JobTitle:       "HR",
					JobDescription: "HR Desc",
					Salary:         "100000",
					CompanyId:      1,
					CategoryId:     2,
				},
				{
					JobTitle:       "SDET",
					JobDescription: "SDET Desc",
					Salary:         "100000",
					CompanyId:      2,
					CategoryId:     1,
				},
			},
			expectedError: nil,
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.FetchAllJobs(ctx, &pb.NoParams{})
		if err != tc.expectedError {
			t.Errorf("Unexpected Error\nExpected: %v\nGot: %v\n", tc.expectedError, err)
		}

		if !reflect.DeepEqual(tc.expectedOutput, got.Jobs) {
			t.Errorf("Mismatch in output:\nExpected: %v\nGot:%v\n", tc.expectedOutput, got)
		}
	}
}
