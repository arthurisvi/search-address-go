package main

import (
	"log"
	"search-address-service/internal/application/services"
	"search-address-service/internal/infrastructure/zipcodeclient/viacep"
	"search-address-service/internal/presentation/web_api/handlers"
)

func main() {
	viacepClient := viacep.NewViaCepClient()

	addressService := services.NewAddressService(viacepClient)

	addressHandlers := handlers.NewAddressHandlers(addressService)

	router := InitializeRoutes(addressHandlers)

	err := router.Run(":8080")

	log.Printf("Starting server on port 8080")

	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
