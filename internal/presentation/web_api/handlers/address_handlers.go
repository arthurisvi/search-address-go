package handlers

import (
	"net/http"
	"via-cep-client/internal/application/services"

	"github.com/gin-gonic/gin"
)

type AddressHandlers struct {
	AdressService *services.AddressService
}

func NewAddressHandlers(zipcodeService *services.AddressService) *AddressHandlers {
	return &AddressHandlers{
		AdressService: zipcodeService,
	}
}

func (z *AddressHandlers) GetAddress(c *gin.Context) {
	zipCode := c.Param("cep")

	address, err := z.AdressService.GetAddressByZipCode(zipCode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CEP não encontrado"})
		return
	}

	c.JSON(http.StatusOK, address)
}
