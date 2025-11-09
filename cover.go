package openlibrary

import (
	"errors"
	"net/http"
	"path"
)

type CoverKind string

const (
	ISBNCover CoverKind = "isbn"
	OCLCCover CoverKind = "oclc"
	LCCNCover CoverKind = "lccn"
	OLIDCover CoverKind = "olid"
	IDCover   CoverKind = "id"
)

type CoverSize string

const (
	CoverSmall  CoverSize = "S"
	CoverMedium CoverSize = "M"
	CoverLarge  CoverSize = "L"
)

type CoverAPI struct {
	openlibraryClient *Client
	value             string
	kind              CoverKind
	size              CoverSize
}

func (c *Client) Cover() *CoverAPI {
	api := &CoverAPI{
		openlibraryClient: c,
		kind:              ISBNCover,
		size:              CoverMedium,
	}
	api.openlibraryClient.changeBaseUrl(coverBaseURL)
	return api
}

func (api *CoverAPI) ISBN(isbn string) *CoverAPI {
	api.kind = ISBNCover
	api.value = isbn
	return api
}

func (api *CoverAPI) OCLC(oclc string) *CoverAPI {
	api.kind = OCLCCover
	api.value = oclc
	return api
}

func (api *CoverAPI) LCCN(lccn string) *CoverAPI {
	api.kind = LCCNCover
	api.value = lccn
	return api
}

func (api *CoverAPI) OLID(olid string) *CoverAPI {
	api.kind = OLIDCover
	api.value = olid
	return api
}

func (api *CoverAPI) ID(id string) *CoverAPI {
	api.kind = IDCover
	api.value = id
	return api
}

func (api *CoverAPI) Small() *CoverAPI {
	api.size = CoverSmall
	return api
}

func (api *CoverAPI) Medium() *CoverAPI {
	api.size = CoverMedium
	return api
}

func (api *CoverAPI) Large() *CoverAPI {
	api.size = CoverLarge
	return api
}

func (api *CoverAPI) Get() (imgBytes []byte, mimeType string, err error) {
	if api.value == "" {
		return imgBytes, mimeType, errors.New("No value")
	}
	endpoint := path.Join("/b", string(api.kind), api.value+"-"+string(api.size)) + ".jpg?default=false"
	res, err := api.openlibraryClient.httpClient.R().Get(endpoint)
	if err != nil {
		return
	}

	if res.StatusCode() != 200 {
		return
	}

	imgBytes = res.Bytes()
	mimeType = http.DetectContentType(imgBytes)
	return
}
