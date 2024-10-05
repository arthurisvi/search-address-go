package services

import (
	"errors"
	"testing"
	"via-cep-client/domain/models"

	"github.com/stretchr/testify/assert"
)

type MockZipCodeClient struct {
	Result *models.AddressModel
	Err    error
}

func (m *MockZipCodeClient) SearchByZipCode(zipCode string) (*models.AddressModel, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Result, nil
}

func TestGetAddressByZipCodeWithErrorFromSearchByZipCode(t *testing.T) {
	mockClient := &MockZipCodeClient{
		Result: nil,
		Err:    errors.New("erro na busca: CEP não existente"),
	}

	zipCodeService := NewZipCodeService(mockClient)

	address, err := zipCodeService.GetAddressByZipCode("99999999")

	assert.NotNil(t, err)
	assert.Equal(t, "erro na busca: CEP não existente", err.Message)
	assert.Nil(t, address)
}

func TestGetAddressByZipCodeWithSuccess(t *testing.T) {
	mockClient := &MockZipCodeClient{
		Result: &models.AddressModel{
			ZipCode: "55026-005",
		},
		Err: nil,
	}

	zipCodeService := NewZipCodeService(mockClient)

	address, err := zipCodeService.GetAddressByZipCode("55026005")

	assert.Nil(t, err)
	assert.Equal(t, "55026-005", address.ZipCode)
}

func TestGetAddressByZipCodeWithErrorByInvalidZipCode(t *testing.T) {
	mockClient := &MockZipCodeClient{
		Result: &models.AddressModel{
			ZipCode: "55026-005",
		},
		Err: nil,
	}

	zipCodeService := NewZipCodeService(mockClient)

	address, err := zipCodeService.GetAddressByZipCode("5502600B")
	address2, err2 := zipCodeService.GetAddressByZipCode("5502600")

	assert.NotNil(t, err)
	assert.Equal(t, "o CEP informado não é válido: não pode conter números", err.Message)
	assert.Equal(t, "o CEP informado não é válido: deve possuir 8 caracteres", err2.Message)
	assert.Nil(t, address)
	assert.Nil(t, address2)
}
