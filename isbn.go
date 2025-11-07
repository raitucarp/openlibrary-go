package openlibrary

import "path"

type ISBNAPI struct {
	openlibraryClient *Client
	isbn              string
}

func (c *Client) ISBN(isbn string) *ISBNAPI {
	api := ISBNAPI{
		openlibraryClient: c,
		isbn:              isbn,
	}
	return &api
}

func (api *ISBNAPI) Get() (resp *Edition, err error) {
	_, err = api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		Get(path.Join("isbn", api.isbn) + ".json")

	if err != nil {
		return
	}
	return
}
