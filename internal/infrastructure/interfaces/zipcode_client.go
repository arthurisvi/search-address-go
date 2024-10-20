package interfaces

import "search-address-service/internal/domain/models"

type ZipCodeClient interface {
	SearchByZipCode(zipCode string) (*models.AddressModel, error)
}
