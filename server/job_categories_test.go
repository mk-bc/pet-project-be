package server

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/mk-bc/pet-project-be/database/data"
	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
	"go.uber.org/mock/gomock"
)

func TestFetchJobCategories(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.FetchJobCategoriesRequest
		mockFunc       func()
		expectedOutput *pb.FetchJobCategoriesResponse
		expectedError  error
	}{
		{
			name:  "Success/Valid",
			input: &pb.FetchJobCategoriesRequest{},
			mockFunc: func() {
				mockDB.EXPECT().FetchJobCategories().Return([]*models.JobCategory{
					{
						CategoryName: "HR Department",
					},
					{
						CategoryName: "Cybersecurity Engineer",
					},
					{
						CategoryName: "Software Engineer",
					},
					{
						CategoryName: "Product Manager",
					},
				}, nil)
			},
			expectedOutput: &pb.FetchJobCategoriesResponse{
				Categories: []*pb.JobCategory{
					{
						CategoryName: "HR Department",
					},
					{
						CategoryName: "Cybersecurity Engineer",
					},
					{
						CategoryName: "Software Engineer",
					},
					{
						CategoryName: "Product Manager",
					},
				},
			},
			expectedError: nil,
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.FetchJobCategories(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}

func TestCreateNewJobCategory(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := data.NewMockDatabase(controller)
	server := &JobPortalServiceServer{Db: mockDB}

	ctx := context.Background()

	testcases := []struct {
		name           string
		input          *pb.CreateNewJobCategoryRequest
		mockFunc       func()
		expectedOutput *pb.CreateNewJobCategoryResponse
		expectedError  error
	}{
		{
			name: "Success/Valid",
			input: &pb.CreateNewJobCategoryRequest{
				Category: &pb.JobCategory{
					CategoryName: "SWE Intern",
				},
			},
			mockFunc: func() {
				mockDB.EXPECT().CreateNewJobCategory(gomock.Any()).Return(nil)
			},
			expectedOutput: &pb.CreateNewJobCategoryResponse{
				Category: &pb.JobCategory{
					CategoryName: "SWE Intern",
				},
			},
			expectedError: nil,
		},
		{
			name: "Incomplete input",
			input: &pb.CreateNewJobCategoryRequest{
				Category: &pb.JobCategory{
					CategoryName: "",
				},
			},
			mockFunc: func() {
				// mockDB.EXPECT().CreateNewJobCategory(gomock.Any()).Return(fmt.Errorf("No name to enter"))
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("No name to insert"),
		},
	}
	for _, tc := range testcases {
		tc.mockFunc()

		got, err := server.CreateNewJobCategory(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expectedError) {
			t.Errorf("Unknown error: Expected: %v, Got: %v", tc.expectedError, err)
		}

		if !reflect.DeepEqual(got, tc.expectedOutput) {
			t.Errorf("Unknown ouput: Expected: %v, Got: %v", tc.expectedOutput, got)
		}
	}
}
