package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lyoncesar/my-commerce/gateways/db/repository"
	"github.com/lyoncesar/my-commerce/internal/domain/customer/usecase"
	"github.com/lyoncesar/my-commerce/internal/domain/customer/usecase/input"
	"github.com/lyoncesar/my-commerce/pkg/config"
)

func main() {
	cfg := config.Initialize()
	fmt.Sprintf("loaded envs: %+v", cfg)

	ctx := context.Background()

	customerRepo := repository.NewCustomerRepository()
	createCustomerUC := usecase.NewCreateCustomerUC(customerRepo)

	createCustomerInput := input.CreateCustomerInput{
		Name:     "Lyon Oliveira",
		Email:    "lyon3@my-commerce.com",
		Document: "12345677800",
	}
	customer, err := createCustomerUC.Execute(ctx, createCustomerInput)
	if err != nil {
		log.Printf("[error] customer creation: ", err)
	}

	log.Printf("Customer created! data: %+v", customer)
}
