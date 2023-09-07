package server

import (
	"context"
	"log"

	pb "github.com/mk-bc/pet-project-be/proto"
)

func (server *JobPortalServiceServer) FetchJobCategories(
	context.Context,
	*pb.FetchJobCategoriesRequest) (*pb.FetchJobCategoriesResponse, error) {
	log.Println("server: fetchJobCategories")

	job_categories, err := server.Db.FetchJobCategories()
	if err != nil {
		log.Fatalf("Error retrieving job cateogories: %v", err)
		return nil, err
	}
	var response []*pb.JobCategory
	for _, cat := range job_categories {
		response = append(response, &pb.JobCategory{
			CategoryId:   uint32(cat.ID),
			CategoryName: cat.CategoryName,
		})
	}
	return &pb.FetchJobCategoriesResponse{
		Categories: response,
	}, nil
}
