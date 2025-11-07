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

type coverAPI struct {
	openlibraryClient *client
	value             string
	kind              CoverKind
	size              CoverSize
}

func (c *client) Cover() *coverAPI {
	api := &coverAPI{
		openlibraryClient: c,
	}
	api.openlibraryClient.changeBaseUrl(coverBaseURL)
	return api
}

func (c *coverAPI) ISBN(isbn string) *coverAPI {
	c.kind = ISBNCover
	c.value = isbn
	return c
}

func (c *coverAPI) OCLC(oclc string) *coverAPI {
	c.kind = OCLCCover
	c.value = oclc
	return c
}

func (c *coverAPI) LCCN(lccn string) *coverAPI {
	c.kind = LCCNCover
	c.value = lccn
	return c
}

func (c *coverAPI) OLID(olid string) *coverAPI {
	c.kind = OLIDCover
	c.value = olid
	return c
}

func (c *coverAPI) ID(id string) *coverAPI {
	c.kind = IDCover
	c.value = id
	return c
}

func (c *coverAPI) Small() *coverAPI {
	c.size = CoverSmall
	return c
}

func (c *coverAPI) Medium() *coverAPI {
	c.size = CoverMedium
	return c
}

func (c *coverAPI) Large() *coverAPI {
	c.size = CoverLarge
	return c
}

func (c *coverAPI) Get() (imgBytes []byte, mimeType string, err error) {
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
