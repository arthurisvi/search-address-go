package httpclient

import "net/http"

type NetHttpClient struct{}

func (c *NetHttpClient) Get(url string) (*http.Response, error) {
	return http.DefaultClient.Get(url)
}

func NewNetHttpClient() *NetHttpClient {
	return &NetHttpClient{}
}
