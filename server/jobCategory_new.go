package server

import (
	"context"
	"log"

	"github.com/mk-bc/pet-project-be/models"
	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) CreateNewJobCategory(
	ctx context.Context,
	req *pb.CreateNewJobCategoryRequest) (*pb.CreateNewJobCategoryResponse, error) {
	log.Println("server: create new job category")
	err := server.Db.CreateNewJobCategory(models.JobCategory{
		CategoryName: req.Category.CategoryName,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateNewJobCategoryResponse{
		Category: req.Category,
	}, nil
}
