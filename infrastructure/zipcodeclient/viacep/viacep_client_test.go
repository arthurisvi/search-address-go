package viacep

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHttpClient struct {
	StatusCode   int
	ResponseBody string
}

func (c *MockHttpClient) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: c.StatusCode,
		Body:       io.NopCloser(bytes.NewBufferString(c.ResponseBody)),
	}, nil
}

type MockHttpClientWithRequestError struct{}

func (c *MockHttpClientWithRequestError) Get(url string) (*http.Response, error) {
	return nil, errors.New("erro ao fazer a requisição HTTP")
}

type MockHttpClientWithReadError struct{}

func (c *MockHttpClientWithReadError) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(&FailingReader{}),
	}, nil
}

type FailingReader struct{}

func (r *FailingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("erro ao ler o corpo da resposta")
}

type MockHttpClientWithUnmarshalError struct{}

func (c *MockHttpClientWithUnmarshalError) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("invalid JSON")),
	}, nil
}

func TestSearchByZipCodeWithSuccess(t *testing.T) {
	t.Run("Should return address when cep is valid", func(t *testing.T) {

		client := NewViaCepClient(&MockHttpClient{
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
				"siafi": "2301"
			}`,
		})

		result, err := client.SearchByZipCode("55026005")

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "55026-005", result.ZipCode)
	})
}

func TestSearchByZipCodeWithStatusNotOk(t *testing.T) {
	t.Run("Should return message with status and API response", func(t *testing.T) {

		client := NewViaCepClient(
			&MockHttpClient{
				StatusCode:   400,
				ResponseBody: `{"erro": "CEP não encontrado"}`,
			},
		)

		result, err := client.SearchByZipCode("00000000")

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "erro ao buscar o CEP: status 400, resposta: {\"erro\": \"CEP não encontrado\"}")
	})
}

func TestSearchByZipCodeWithStatusOkButHasError(t *testing.T) {
	t.Run("Should return error zipcode not found", func(t *testing.T) {

		client := NewViaCepClient(
			&MockHttpClient{
				StatusCode:   200,
				ResponseBody: `{"erro": "true"}`,
			},
		)

		result, err := client.SearchByZipCode("00000000")

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "erro ao buscar o CEP: CEP não encontrado")
	})
}

func TestSearchByZipCodeWithRequestError(t *testing.T) {
	t.Run("Should return error when httpClient.Get fails", func(t *testing.T) {
		client := NewViaCepClient(
			&MockHttpClientWithRequestError{},
		)

		result, err := client.SearchByZipCode("55026005")

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "erro ao fazer a requisição HTTP", err.Error())
	})
}

func TestSearchByZipCodeWithReadError(t *testing.T) {
	t.Run("Should return error when io.ReadAll fails", func(t *testing.T) {
		client := NewViaCepClient(
			&MockHttpClientWithReadError{},
		)

		result, err := client.SearchByZipCode("55026005")

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "erro ao ler o corpo da resposta", err.Error())
	})
}

func TestSearchByZipCodeWithUnmarshalError(t *testing.T) {
	t.Run("Should return error when json.Unmarshal fails", func(t *testing.T) {
		client := NewViaCepClient(
			&MockHttpClientWithUnmarshalError{},
		)

		result, err := client.SearchByZipCode("55026005")

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "invalid character")
	})
}
