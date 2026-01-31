package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/lyoncesar/my-commerce/gateways/db/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) Create(ctx context.Context, customerEntity model.Customer) (*model.Customer, error) {
	args := m.Called(ctx, customerEntity)
	return args.Get(0).(*model.Customer), args.Get(1).(error)
}

func CreateCustomerUCTEst_Execute(t *testing.T) {

	ctx := context.Background()
	repository := &CustomerRepositoryMock{}
	customerUUID := uuid.New()
	anyCustomer := model.Customer{
		Name:     "Road Runner Bip Bip",
		Email:    "road_runner@acme.com",
		Document: "12345678800",
	}
	wantCustomer := model.Customer{
		UUID:     customerUUID,
		Name:     "Road Runner Bip Bip",
		Email:    "road_runner@acme.com",
		Document: "12345678800",
	}

	t.Run("Create a customer when receives valid params", func(t *testing.T) {

		repository.On("Create", ctx, anyCustomer).Return(wantCustomer, nil)

		uc := NewCreateCustomerUC(repository)
		got, err := uc.repository.Create(ctx, anyCustomer)

		if assert.NoError(t, err) {
			assert.Equal(t, got, wantCustomer)
		}
	})

	t.Run("Should not create a customer when repository return an error", func(t *testing.T) {
		repository.On("Create", ctx, anyCustomer).Return(nil, assert.AnError)

		uc := NewCreateCustomerUC(repository)
		got, err := uc.repository.Create(ctx, anyCustomer)

		assert.Error(t, err)
		assert.Nil(t, got)
	})
}
