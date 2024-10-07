package viacep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"via-cep-client/application/interfaces"
	"via-cep-client/domain/models"
	httpclient "via-cep-client/infrastructure/http"
)

type viaCepClient struct {
	httpClient interfaces.HttpClient
}

func (c *viaCepClient) SearchByZipCode(zipCode string) (*models.AddressModel, error) {
	r, err := c.httpClient.Get(("https://viacep.com.br/ws/" + zipCode + "/json/"))

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	response, readError := io.ReadAll(r.Body)

	if readError != nil {
		return nil, readError
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao buscar o CEP: status %d, resposta: %s", r.StatusCode, string(response))
	}

	responseDTO := ViaCepResponseDTO{}
	err = json.Unmarshal(response, &responseDTO)

	if err != nil {
		return nil, err
	}

	if responseDTO.Erro != "" {
		return nil, fmt.Errorf("erro ao buscar o CEP: CEP não encontrado")
	}

	return responseDTO.ToDomain(), nil
}

func NewViaCepClient() *viaCepClient {
	return &viaCepClient{httpClient: httpclient.NewNetHttpClient()}
}
