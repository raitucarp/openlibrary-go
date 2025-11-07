package openlibrary

import "resty.dev/v3"

const baseURL = "https://openlibrary.org"
const coverBaseURL = "https://covers.openlibrary.org/b"
const userAgent = "openlibrary-go/1.0 (Golang OpenLibrary Client)"

type client struct {
	httpClient *resty.Client
	userAgent  string
}

func NewClient() *client {
	httpClient := resty.New()
	httpClient.SetBaseURL(baseURL)
	httpClient.SetHeader("User-Agent", userAgent)

	c := client{httpClient: httpClient}
	return &c
}

func (c *client) SetUserAgent(useragent string) {
	c.httpClient.SetHeader("User-Agent", userAgent)
}

func (c *client) changeBaseUrl(baseUrl string) {
	c.httpClient.SetBaseURL(baseUrl)
}
