package openlibrary

import "resty.dev/v3"

const baseURL = "https://openlibrary.org"
const coverBaseURL = "https://covers.openlibrary.org"
const userAgent = "openlibrary-go/1.0 (Golang OpenLibrary Client)"

type Client struct {
	httpClient *resty.Client
	userAgent  string
}

func NewClient() *Client {
	httpClient := resty.New()
	httpClient.SetBaseURL(baseURL)
	httpClient.SetHeader("User-Agent", userAgent)

	c := Client{httpClient: httpClient}
	return &c
}

func (c *Client) SetUserAgent(useragent string) {
	c.httpClient.SetHeader("User-Agent", userAgent)
}

func (c *Client) changeBaseUrl(baseUrl string) {
	c.httpClient.SetBaseURL(baseUrl)
}
