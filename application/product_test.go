package application_test 

import (
	"github.com/glaydsonp/go-hexagonal/application"
	"testing"
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
	
	product.Price = -10
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test"
	product.Status = application.DISABLED
	product.Price = 0
	
	err := product.Disable()
	require.Nil(t, err)
	
	product.Price = -10
	err = product.Disable()
	require.Nil(t, err)
	
	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be less than or equal to zero in order to disable the product", err.Error())
	
}