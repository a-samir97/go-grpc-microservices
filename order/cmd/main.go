package main

import (
	"log"

	"Github.com/a-samir97/microservices/order/config"
	"Github.com/a-samir97/microservices/order/internal/adapters/db"
	"Github.com/a-samir97/microservices/order/internal/adapters/grpc"
	"Github.com/a-samir97/microservices/order/internal/adapters/payment"
	"Github.com/a-samir97/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())

	if err != nil {
		log.Fatalf("Failed to connect to database. Error : %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceURL())
	if err != nil {
		log.Fatalf("Failed to inialize payment stub. Error : %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)

	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
