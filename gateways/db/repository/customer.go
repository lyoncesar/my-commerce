package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lyoncesar/my-commerce/gateways/db/model"
	"github.com/lyoncesar/my-commerce/pkg/client/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient interface {
	Open(dbConfig db.DBConfig) (*gorm.DB, error)
}

type customerRepository struct{}

func NewCustomerRepository() *customerRepository {
	return &customerRepository{}
}

func (r *customerRepository) Create(ctx context.Context, customer model.Customer) (*model.Customer, error) {
	db := connectDB()
	db.AutoMigrate(model.Customer{})

	newUUID, err := uuid.NewV7()
	if err != nil {
		newUUID = uuid.New()
	}

	customer.UUID = newUUID

	err = gorm.G[model.Customer](db).Create(ctx, &customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=GMT",
		"localhost",
		"my-commerce-admin",
		"password",
		"my-commerce-db",
		"5432",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return db
}
