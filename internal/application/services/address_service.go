package services

import (
	"errors"
	"regexp"
	"via-cep-client/internal/domain/models"
	"via-cep-client/internal/infrastructure/interfaces"
)

type AddressService struct {
	client interfaces.ZipCodeClient
}

type ErrorResponseDTO struct {
	Message string
}

func (s *AddressService) GetAddressByZipCode(zipCode string) (*models.AddressModel, *ErrorResponseDTO) {
	hasError := s.validZipCode(zipCode)
	if hasError != nil {
		return nil, &ErrorResponseDTO{
			Message: "o CEP informado não é válido: " + hasError.Error(),
		}
	}

	address, err := s.client.SearchByZipCode(zipCode)

	if err != nil {
		return nil, &ErrorResponseDTO{
			Message: err.Error(),
		}
	}

	return address, nil
}

func (s *AddressService) validZipCode(zipCode string) error {
	if len(zipCode) != 8 {
		return errors.New("deve possuir 8 caracteres")
	}

	numberRegex := regexp.MustCompile(`^\d+$`)
	if !numberRegex.MatchString(zipCode) {
		return errors.New("não pode conter números")
	}

	return nil
}

func NewAddressService(client interfaces.ZipCodeClient) *AddressService {
	return &AddressService{
		client: client,
	}
}
