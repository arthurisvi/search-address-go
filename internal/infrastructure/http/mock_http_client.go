package httpclient

import (
	"bytes"
	"errors"
	"io"
	"net/http"
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
