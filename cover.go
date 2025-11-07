package openlibrary

import (
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
	}
	api.openlibraryClient.changeBaseUrl(coverBaseURL)
	return api
}

func (c *CoverAPI) ISBN(isbn string) *CoverAPI {
	c.kind = ISBNCover
	c.value = isbn
	return c
}

func (c *CoverAPI) OCLC(oclc string) *CoverAPI {
	c.kind = OCLCCover
	c.value = oclc
	return c
}

func (c *CoverAPI) LCCN(lccn string) *CoverAPI {
	c.kind = LCCNCover
	c.value = lccn
	return c
}

func (c *CoverAPI) OLID(olid string) *CoverAPI {
	c.kind = OLIDCover
	c.value = olid
	return c
}

func (c *CoverAPI) ID(id string) *CoverAPI {
	c.kind = IDCover
	c.value = id
	return c
}

func (c *CoverAPI) Small() *CoverAPI {
	c.size = CoverSmall
	return c
}

func (c *CoverAPI) Medium() *CoverAPI {
	c.size = CoverMedium
	return c
}

func (c *CoverAPI) Large() *CoverAPI {
	c.size = CoverLarge
	return c
}

func (c *CoverAPI) Get() (imgBytes []byte, mimeType string, err error) {
	endpoint := path.Join(string(c.kind), c.value+"-"+string(c.size)) + ".jpg?default=false"
	res, err := c.openlibraryClient.httpClient.R().Get(endpoint)
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
