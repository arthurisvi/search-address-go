package httpclient

import "net/http"

type netHttpClient struct{}

func (c *netHttpClient) Get(url string) (*http.Response, error) {
	return http.DefaultClient.Get(url)
}

func NewNetHttpClient() *netHttpClient {
	return &netHttpClient{}
}
