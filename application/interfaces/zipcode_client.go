package interfaces

import "via-cep-client/domain/models"

type ZipCodeClient interface {
	SearchByZipCode(zipCode string) (*models.AddressModel, error)
}
