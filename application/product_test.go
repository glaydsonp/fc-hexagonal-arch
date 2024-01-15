package application_test

import (
	"testing"

	"github.com/glaydsonp/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test"
	product.Status = application.DISABLED

	product.Price = 10
	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be less than or equal to zero in order to disable the product", err.Error())
}

func TestProduct_IsVakid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Test"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater than zero", err.Error())
}
