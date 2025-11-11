package openlibrary

import (
	"errors"
	"path"
)

type ReadIDType string

const (
	ReadISBN ReadIDType = "isbn"
	ReadLCCN ReadIDType = "lccn"
	ReadOCLC ReadIDType = "oclc"
	ReadOLID ReadIDType = "olid"
)

type ReadAPI struct {
	openlibraryClient *Client
	idType            ReadIDType
	idValue           string
}

func (c *Client) Read() *ReadAPI {
	api := ReadAPI{
		openlibraryClient: c,
	}
	return &api
}

func (result *SearchResult) ReadByISBN() (readAPI *ReadAPI, err error) {
	if len(result.ISBN) <= 0 {
		err = errors.New("No ISBN")
		return
	}

	readAPI = &ReadAPI{
		openlibraryClient: NewClient(),
	}

	for _, isbn := range result.ISBN {
		if readAPI.idValue == "" {
			readAPI.ISBN(isbn)
		}
	}

	return
}

func (result *SearchResult) ReadByLCCN() (readAPI *ReadAPI, err error) {
	if len(result.LCCN) <= 0 {
		err = errors.New("No LCCN")
		return
	}

	readAPI = &ReadAPI{
		openlibraryClient: NewClient(),
	}

	for _, lccn := range result.LCCN {
		if readAPI.idValue == "" {
			readAPI.LCCN(lccn)
		}
	}

	return
}

func (result *SearchResult) ReadByOCLC() (readAPI *ReadAPI, err error) {
	if len(result.OCLC) <= 0 {
		err = errors.New("No OCLC")
		return
	}

	readAPI = &ReadAPI{
		openlibraryClient: NewClient(),
	}

	for _, oclc := range result.OCLC {
		if readAPI.idValue == "" {
			readAPI.OCLC(oclc)
		}
	}

	return
}

func (result *SearchResult) Read() (readAPI *ReadAPI, err error) {
	if result.CoverEditionKey == "" {
		return
	}

	readAPI = &ReadAPI{
		openlibraryClient: NewClient(),
	}

	readAPI.OLID(result.CoverEditionKey)

	return
}

func (edition *Edition) ReadByISBN() (readAPI *ReadAPI) {
	if len(edition.ISBN10) <= 0 || len(edition.ISBN13) <= 0 {
		return
	}

	readAPI = &ReadAPI{
		openlibraryClient: NewClient(),
	}

	if len(edition.ISBN10) > 0 {
		for _, isbn := range edition.ISBN10 {
			if readAPI.idValue == "" {
				readAPI.ISBN(isbn)
			}
		}
	}

	if len(edition.ISBN13) > 0 {
		for _, isbn := range edition.ISBN13 {
			if readAPI.idValue == "" {
				readAPI.ISBN(isbn)
			}
		}
	}

	return
}

func (edition *Edition) ReadByLCCN() (readAPI *ReadAPI) {
	if len(edition.LCCN) <= 0 {
		return
	}

	readAPI = &ReadAPI{
		openlibraryClient: NewClient(),
	}

	for _, lccn := range edition.LCCN {
		if readAPI.idValue == "" {
			readAPI.LCCN(lccn)
		}
	}

	return
}

func (edition *Edition) ReadByOCLC() (readAPI *ReadAPI) {
	if len(edition.OCLCNumbers) <= 0 {
		return
	}

	readAPI = &ReadAPI{
		openlibraryClient: NewClient(),
	}

	for _, oclc := range edition.OCLCNumbers {
		if readAPI.idValue == "" {
			readAPI.OCLC(oclc)
		}
	}

	return
}

func (edition *Edition) Read() (readAPI *ReadAPI) {

	readAPI = &ReadAPI{
		openlibraryClient: NewClient(),
	}

	readAPI.OLID(edition.Key)

	return
}

func (api *ReadAPI) ISBN(isbn string) *ReadAPI {
	api.idType = ReadISBN
	api.idValue = isbn
	return api
}

func (api *ReadAPI) LCCN(lccn string) *ReadAPI {
	api.idType = ReadLCCN
	api.idValue = lccn
	return api
}

func (api *ReadAPI) OCLC(oclc string) *ReadAPI {
	api.idType = ReadOCLC
	api.idValue = oclc
	return api
}

func (api *ReadAPI) OLID(olid string) *ReadAPI {
	api.idType = ReadOLID
	api.idValue = olid
	return api
}

type Cover struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type ReadItems struct {
	Enumcron    bool   `json:"enumcron,omitempty"`
	Match       string `json:"match,omitempty"`
	Status      string `json:"status,omitempty"`
	FromRecord  string `json:"fromRecord,omitempty"`
	OlEditionID string `json:"ol-edition-id,omitempty"`
	OlWorkID    string `json:"ol-work-id,omitempty"`
	PublishDate string `json:"publishDate,omitempty"`
	Contributor string `json:"contributor,omitempty"`
	ItemURL     string `json:"itemURL,omitempty"`
	Cover       *Cover `json:"cover,omitempty"`
}

type ReadDetails struct {
	BibKey       string   `json:"bib_key,omitempty"`
	InfoURL      string   `json:"info_url,omitempty"`
	Preview      string   `json:"preview,omitempty"`
	PreviewURL   string   `json:"preview_url,omitempty"`
	ThumbnailURL string   `json:"thumbnail_url,omitempty"`
	Details      *Edition `json:"details,omitempty"`
}

type ReadData struct {
	URL             string      `json:"url"`
	Key             string      `json:"key"`
	Title           string      `json:"title"`
	Subtitle        string      `json:"subtitle"`
	NumberOfPages   int         `json:"number_of_pages,omitempty"`
	Weight          string      `json:"weight,omitempty"`
	Identifiers     Identifiers `json:"identifiers"`
	Classifications struct {
		LCClassifications []string `json:"lc_classifications"`
		DeweyDecimalClass []string `json:"dewey_decimal_class"`
	} `json:"classifications"`

	PublishPlaces []struct {
		Name string `json:"name"`
	} `json:"publish_places"`
	PublishDate string    `json:"publish_date"`
	Subjects    []KeyName `json:"subjects,omitempty"`
	Places      []KeyName `json:"subject_places,omitempty"`
	People      []KeyName `json:"subject_people,omitempty"`
	Times       []KeyName `json:"subject_times,omitempty"`
	Authors     []KeyName `json:"authors,omitempty"`
	Publishers  []KeyName `json:"publishers,omitempty"`
	Languages   []KeyName `json:"languages,omitempty"`
	Ebooks      []struct {
		PreviewURL   string `json:"preview_url,omitempty"`
		Availability string `json:"availability,omitempty"`
		Formats      any    `json:"formats,omitempty"`
	} `json:"ebooks,omitempty"`
	Cover Cover `json:"cover"`
}

type ReadRecord struct {
	Isbns        []string    `json:"isbns"`
	Issns        []any       `json:"issns"`
	Lccns        []string    `json:"lccns"`
	Oclcs        []string    `json:"oclcs"`
	Olids        []string    `json:"olids"`
	PublishDates []string    `json:"publishDates"`
	RecordURL    string      `json:"recordURL"`
	Data         ReadData    `json:"data"`
	Details      ReadDetails `json:"details"`
}

type ReadResponse struct {
	Records map[string]ReadRecord `json:"records"`
	Items   []ReadItems           `json:"items"`
}

func (api *ReadAPI) Get() (resp *ReadResponse, err error) {
	endpoint := path.Join("/api", "volumes", "brief", string(api.idType), api.idValue+".json")

	_, err = api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		Get(endpoint)

	if err != nil {
		return
	}
	return
}
