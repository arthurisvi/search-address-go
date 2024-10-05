package opencep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"via-cep-client/application/interfaces"
	"via-cep-client/domain/models"
)

type OpenCepClient struct {
	httpClient interfaces.HttpClient
}

func (c *OpenCepClient) SearchByZipCode(zipCode string) (*models.AddressModel, error) {
	r, err := c.httpClient.Get(("https://opencep.com/v1/" + zipCode + ".json"))

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	response, readError := io.ReadAll(r.Body)

	if readError != nil {
		return nil, readError
	}

	if r.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("erro ao buscar o CEP: cep não existente")
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao buscar o CEP: status %d, resposta: %s", r.StatusCode, string(response))
	}

	responseDTO := OpenCepResponseDTO{}
	err = json.Unmarshal(response, &responseDTO)

	if err != nil {
		return nil, err
	}

	return responseDTO.ToDomain(), nil
}

func NewOpenCepClient(c interfaces.HttpClient) *OpenCepClient {
	return &OpenCepClient{httpClient: c}
}
