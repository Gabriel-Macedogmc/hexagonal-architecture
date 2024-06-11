package application_test

import (
	"testing"

	"github.com/Gabriel-Macedogmc/hexagonal-architecture/application"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product A"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "the price must be greater than 0 to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product A"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, "the price must be 0", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewString()
	product.Name = "Product A"
	product.Status = application.ENABLED
	product.Price = 10

	isValid, err := product.IsValid()
	require.Equal(t, isValid, true)
	require.Nil(t, err)

	product.Status = "INVALID"
	isValid, err = product.IsValid()
	require.Equal(t, isValid, false)
	require.Equal(t, "the status must be DISABLED or ENABLED", err.Error())

	product.Status = application.ENABLED
	isValid, err = product.IsValid()
	require.Equal(t, isValid, true)
	require.Nil(t, err)

	product.Price = -1
	isValid, err = product.IsValid()
	require.Equal(t, isValid, false)
	require.Equal(t, "the price must be greater than 0 to enable the product", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewString()
	product.Name = "Product A"
	product.Status = application.ENABLED
	product.Price = 10

	id := product.GetID()

	require.Equal(t, product.ID, id)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewString()
	product.Name = "Product A"
	product.Status = application.ENABLED
	product.Price = 10

	name := product.GetName()

	require.Equal(t, "Product A", name)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewString()
	product.Name = "Product A"
	product.Status = application.ENABLED
	product.Price = 10

	status := product.GetStatus()

	require.Equal(t, application.ENABLED, status)
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewString()
	product.Name = "Product A"
	product.Status = application.ENABLED
	product.Price = 10

	price := product.GetPrice()

	require.Equal(t, float64(10), price)
}
