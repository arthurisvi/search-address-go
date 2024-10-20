package interfaces

import "net/http"

type HttpClient interface {
	// TO DO: melhorar para não utilizar o http response, visto que o sentido é ser desacoplado do package net/http
	Get(url string) (*http.Response, error)
}
