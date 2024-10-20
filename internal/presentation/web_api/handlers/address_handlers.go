package handlers

import (
	"net/http"
	"via-cep-client/internal/application/services"

	"github.com/gin-gonic/gin"
)

type AddressHandlers struct {
	ZipcodeService *services.ZipCodeService
}

func NewAddressHandlers(zipcodeService *services.ZipCodeService) *AddressHandlers {
	return &AddressHandlers{
		ZipcodeService: zipcodeService,
	}
}

func (z *AddressHandlers) GetAddress(c *gin.Context) {
	zipCode := c.Param("cep")

	address, err := z.ZipcodeService.GetAddressByZipCode(zipCode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CEP não encontrado"})
		return
	}

	c.JSON(http.StatusOK, address)
}
