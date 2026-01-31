package usecase

import (
	"context"

	"github.com/lyoncesar/my-commerce/gateways/db/model"
	"github.com/lyoncesar/my-commerce/internal/domain/customer/usecase/input"
	"github.com/lyoncesar/my-commerce/internal/domain/customer/usecase/output"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer model.Customer) (*model.Customer, error)
}

type createCustomerUC struct {
	repository CustomerRepository
}

func NewCreateCustomerUC(repository CustomerRepository) *createCustomerUC {
	return &createCustomerUC{
		repository: repository,
	}
}

func (uc *createCustomerUC) Execute(ctx context.Context, customerInput input.CreateCustomerInput) (*output.CreateCustomerOutput, error) {
	customer := model.Customer{
		Name:     customerInput.Name,
		Email:    customerInput.Email,
		Document: customerInput.Document,
	}

	result, err := uc.repository.Create(ctx, customer)
	if err != nil {
		return nil, err
	}

	customerOutput := output.CreateCustomerOutput{
		ID:       result.ID,
		UUID:     result.UUID.String(),
		Name:     result.Name,
		Email:    result.Email,
		Document: result.Document,
	}

	return &customerOutput, nil
}
