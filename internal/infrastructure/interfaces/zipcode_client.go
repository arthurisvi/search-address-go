package interfaces

import "via-cep-client/internal/domain/models"

type ZipCodeClient interface {
	SearchByZipCode(zipCode string) (*models.AddressModel, error)
}
