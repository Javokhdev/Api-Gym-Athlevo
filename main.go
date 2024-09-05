package main

import (
	"fmt"
	"log"

	"api/api"
	"api/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cnf := config.Load()
	authservice, err := grpc.NewClient(cnf.AUTH_PORT, grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to authservice service %v", err)
	}
	defer authservice.Close()

	companyService, err := grpc.NewClient("gym"+cnf.BUDGETING_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to learning service service %v", err)
	}

	router := api.NewGin(companyService)

	fmt.Println("API Gateway running on http://localhost:8050")
	if err := router.Run(":8050"); err != nil {
		log.Fatalf("Failed to connect to gin engine: %v", err)
	}

}
