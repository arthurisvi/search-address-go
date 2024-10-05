package main

import (
	"fmt"
	"via-cep-client/application/services"
	nethttpclient "via-cep-client/infrastructure/http"
	"via-cep-client/infrastructure/zipcodeclient/viacep"
)

func main() {
	client := viacep.NewViaCepClient(nethttpclient.NewNetHttpClient())
	zipCodeService := services.NewZipCodeService(client)

	response, err := zipCodeService.GetAddressByZipCode("55026005")
	//response, err := zipCodeService.GetAddressByZipCode("12345678")

	if err != nil {
		fmt.Printf("%s", err.Message)
	} else {
		fmt.Printf("Endereço: %s", response)
	}
}
