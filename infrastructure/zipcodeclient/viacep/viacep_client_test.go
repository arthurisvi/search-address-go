package viacep

import (
	"testing"
	httpclient "via-cep-client/infrastructure/http"

	"github.com/stretchr/testify/assert"
)

func TestSearchByZipCodeWithSuccess(t *testing.T) {
	t.Run("Should return address when cep is valid", func(t *testing.T) {

		mockClient := httpclient.MockHttpClient{
			StatusCode: 200,
			ResponseBody: `{
				"cep": "55026-005",
				"logradouro": "Rua de Teste",
				"complemento": "",
				"bairro": "Bairro Teste",
				"localidade": "Cidade Teste",
				"uf": "PE",
				"ibge": "2607901",
				"gia": "",
				"ddd": "81",
				"siafi": "2301"}`,
		}
		client := viaCepClient{
			httpClient: &mockClient,
		}

		result, err := client.SearchByZipCode("55026005")

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "55026-005", result.ZipCode)
	})
}

func TestSearchByZipCodeWithStatusNotOk(t *testing.T) {
	t.Run("Should return message with status and API response", func(t *testing.T) {

		client := viaCepClient{
			httpClient: &httpclient.MockHttpClient{
				StatusCode:   400,
				ResponseBody: `{"erro": "CEP não encontrado"}`,
			},
		}

		result, err := client.SearchByZipCode("00000000")

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "erro ao buscar o CEP: status 400, resposta: {\"erro\": \"CEP não encontrado\"}")
	})
}

func TestSearchByZipCodeWithStatusOkButHasError(t *testing.T) {
	t.Run("Should return error zipcode not found", func(t *testing.T) {

		client := viaCepClient{
			httpClient: &httpclient.MockHttpClient{
				StatusCode:   200,
				ResponseBody: `{"erro": "true"}`,
			},
		}

		result, err := client.SearchByZipCode("00000000")

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "erro ao buscar o CEP: CEP não encontrado")
	})
}

func TestSearchByZipCodeWithRequestError(t *testing.T) {
	t.Run("Should return error when httpClient.Get fails", func(t *testing.T) {
		client := viaCepClient{
			httpClient: &httpclient.MockHttpClientWithRequestError{},
		}

		result, err := client.SearchByZipCode("55026005")

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "erro ao fazer a requisição HTTP", err.Error())
	})
}

func TestSearchByZipCodeWithReadError(t *testing.T) {
	t.Run("Should return error when io.ReadAll fails", func(t *testing.T) {
		client := viaCepClient{
			httpClient: &httpclient.MockHttpClientWithReadError{},
		}

		result, err := client.SearchByZipCode("55026005")

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "erro ao ler o corpo da resposta", err.Error())
	})
}

func TestSearchByZipCodeWithUnmarshalError(t *testing.T) {
	t.Run("Should return error when json.Unmarshal fails", func(t *testing.T) {
		client := viaCepClient{
			httpClient: &httpclient.MockHttpClientWithUnmarshalError{},
		}

		result, err := client.SearchByZipCode("55026005")

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "invalid character")
	})
}
