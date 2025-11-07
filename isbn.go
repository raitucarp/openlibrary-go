package openlibrary

import "path"

type isbnAPI struct {
	openlibraryClient *client
	isbn              string
}

func (c *client) ISBN(isbn string) *isbnAPI {
	api := isbnAPI{
		openlibraryClient: c,
		isbn:              isbn,
	}
	return &api
}

func (api *isbnAPI) Get() (resp *Edition, err error) {
	_, err = api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		Get(path.Join("isbn", api.isbn) + ".json")

	if err != nil {
		return
	}
	return
}
