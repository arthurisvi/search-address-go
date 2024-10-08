package main

import (
	"fmt"
	"via-cep-client/application/services"
	"via-cep-client/infrastructure/zipcodeclient/opencep"
	"via-cep-client/infrastructure/zipcodeclient/viacep"
)

func main() {
	viacepClient := viacep.NewViaCepClient()
	opencepClient := opencep.NewOpenCepClient()

	zipCodeServiceByViacep := services.NewZipCodeService(viacepClient)
	responseByViacep, errByViacep := zipCodeServiceByViacep.GetAddressByZipCode("55026005")
	//response, err := zipCodeService.GetAddressByZipCode("12345678")

	if errByViacep != nil {
		fmt.Printf("[Viacep] %s", errByViacep.Message)
	} else {
		fmt.Printf("[Viacep] Endereço: %s", responseByViacep)
	}

	zipCodeServiceByOpencep := services.NewZipCodeService(opencepClient)
	responseByOpencep, errByOpencep := zipCodeServiceByOpencep.GetAddressByZipCode("55026005")
	//response, err := zipCodeService.GetAddressByZipCode("00000000")
	fmt.Printf("\n")

	if errByOpencep != nil {
		fmt.Printf("[Opencep] %s", errByOpencep.Message)
	} else {
		fmt.Printf("[Opencep] Endereço: %s", responseByOpencep)
	}

}
