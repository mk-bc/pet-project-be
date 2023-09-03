package main

import (
	"log"

	"github.com/mk-bc/pet-project-be/client"
	pb "github.com/mk-bc/pet-project-be/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = "8080"

func main() {
	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error establishing client connection: %v", err)
	}

	grpcClient := pb.NewJobPortalServiceClient(conn)

	// client.ListAllJobs(grpcClient)
	// client.FetchCompanyData(grpcClient, 1) //companyID
	// client.FetchUserData(grpcClient, 1)
	// client.FetchJobCategories(grpcClient)
	// client.FetchJobsByCompanyID(grpcClient, 1)
	// client.FetchJobsByCategoryID(grpcClient, 2)
	client.FetchCompanyJobsByCategory(grpcClient, 1, 2)
}
