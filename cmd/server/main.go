package main

import (
	"log"
	"net"

	"github.com/mk-bc/pet-project-be/database/config"
	"github.com/mk-bc/pet-project-be/database/data"
	pb "github.com/mk-bc/pet-project-be/proto"
	"github.com/mk-bc/pet-project-be/server"
	"google.golang.org/grpc"
)

const port = ":8080"

func main() {

	db := config.DBSetup()
	// port := flag.String("port", "8080", "opening-port")
	// flag.Parse()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	grpcServer := grpc.NewServer()

	jobPortalServiceServer := server.JobPortalServiceServer{Db: &data.DBClient{Db: db}}
	pb.RegisterJobPortalServiceServer(grpcServer, &jobPortalServiceServer)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server on port 8080: %v", err)
	}

	log.Printf("Started server on port 8080")
}
