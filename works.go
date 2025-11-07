package openlibrary

import (
	"encoding/json"
	"path"
)

type WorksAPI struct {
	openlibraryClient *Client
	key               string
}

type WorkResponse struct {
	Key               string            `json:"key"`
	Title             string            `json:"title"`
	Subtitle          string            `json:"subtitle,omitempty"`
	Type              WorkType          `json:"type"`
	Authors           []AuthorRole      `json:"authors,omitempty"`
	Covers            []int64           `json:"covers,omitempty"`
	Links             []Link            `json:"links,omitempty"`
	ID                *int64            `json:"id,omitempty"`
	LCClassifications []string          `json:"lc_classifications,omitempty"`
	Subjects          []string          `json:"subjects,omitempty"`
	FirstPublishDate  string            `json:"first_publish_date,omitempty"`
	Description       string            `json:"description,omitempty"`
	Notes             string            `json:"notes,omitempty"`
	Revision          int64             `json:"revision"`
	LatestRevision    int64             `json:"latest_revision,omitempty"`
	Created           *InternalDateTime `json:"created,omitempty"`
	LastModified      *InternalDateTime `json:"last_modified"`
}

type WorkType struct {
	Key string `json:"key"`
}

type AuthorRole struct {
	Type   AuthorRoleType `json:"type"`
	Author Author         `json:"author"`
	Role   *string        `json:"role,omitempty"`
	As     *string        `json:"as,omitempty"`
}

type AuthorRoleType struct {
	Key string `json:"key"`
}

type Author struct {
	Key string `json:"key"`
}

type TextBlock struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Raw   string `json:"-"`
}

func (tb *TextBlock) UnmarshalJSON(data []byte) error {

	var rawString string
	if err := json.Unmarshal(data, &rawString); err == nil {
		tb.Raw = rawString
		return nil
	}

	var textBlockStruct struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	if err := json.Unmarshal(data, &textBlockStruct); err != nil {
		return err
	}

	tb.Type = textBlockStruct.Type
	tb.Value = textBlockStruct.Value
	return nil
}

type InternalDateTime struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Link struct {
	URL   string    `json:"url"`
	Title string    `json:"title"`
	Type  *LinkType `json:"type,omitempty"`
}

type LinkType struct {
	Key string `json:"key"`
}

type WorkKey string
type AuthorKey string
type EditionKey string

type LanguageCode string
type LcClassification string
type PublishCountry string

func (c *Client) Works(key string) *WorksAPI {
	api := WorksAPI{
		openlibraryClient: c,
		key:               key,
	}
	return &api
}

func (api *WorksAPI) Get() (resp *WorkResponse, err error) {
	_, err = api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		Get(api.key + ".json")

	if err != nil {
		return
	}
	return
}

type Rating struct {
	Summary struct {
		Average  int `json:"average,omitempty"`
		Count    int `json:"count,omitempty"`
		Sortable int `json:"sortable,omitempty"`
	} `json:"summary"`

	Counts struct {
		One   int `json:"1,omitempty"`
		Two   int `json:"2,omitempty"`
		Three int `json:"3,omitempty"`
		Four  int `json:"4,omitempty"`
		Five  int `json:"5,omitempty"`
	} `json:"counts"`
}

func (api *WorksAPI) Ratings() (resp *Rating, err error) {
	_, err = api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		Get(path.Join(api.key, "ratings") + ".json")

	if err != nil {
		return
	}
	return
}

type Bookshelve struct {
	Counts struct {
		WantToRead       int `json:"want_to_read,omitempty"`
		CurrentlyReading int `json:"currently_reading,omitempty"`
		AlreadyRead      int `json:"already_read,omitempty"`
	} `json:"counts"`
}

func (api *WorksAPI) Bookshelves() (resp *Bookshelve, err error) {
	_, err = api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		Get(path.Join(api.key, "bookshelves") + ".json")

	if err != nil {
		return
	}
	return
}
