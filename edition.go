package openlibrary

import (
	"path"
)

type Edition struct {
	Key                string            `json:"key"`
	Title              string            `json:"title"`
	Subtitle           *string           `json:"subtitle,omitempty"`
	Type               EditionType       `json:"type"`
	Authors            []Author          `json:"authors,omitempty"`
	Works              []WorkReference   `json:"works"`
	Identifiers        map[string]any    `json:"identifiers,omitempty"`
	ISBN10             []string          `json:"isbn_10,omitempty"`
	ISBN13             []string          `json:"isbn_13,omitempty"`
	LCCN               []string          `json:"lccn,omitempty"`
	OcaId              string            `json:"ocaid,omitempty"`
	OCLCNumbers        []string          `json:"oclc_numbers,omitempty"`
	LocalID            []string          `json:"local_id,omitempty"`
	Covers             []int64           `json:"covers,omitempty"`
	Links              []Link            `json:"links,omitempty"`
	Languages          []Language        `json:"languages,omitempty"`
	TranslatedFrom     []Language        `json:"translated_from,omitempty"`
	TranslationOf      string            `json:"translation_of,omitempty"`
	ByStatement        string            `json:"by_statement,omitempty"`
	Weight             string            `json:"weight,omitempty"`
	EditionName        string            `json:"edition_name,omitempty"`
	NumberOfPages      int64             `json:"number_of_pages,omitempty"`
	Pagination         string            `json:"pagination,omitempty"`
	PhysicalDimensions string            `json:"physical_dimensions,omitempty"`
	PhysicalFormat     string            `json:"physical_format,omitempty"`
	CopyrightDate      string            `json:"copyright_date,omitempty"`
	PublishCountry     string            `json:"publish_country,omitempty"` // Pattern: "^[a-z]{2,3}$"
	PublishDate        string            `json:"publish_date,omitempty"`
	PublishPlaces      []string          `json:"publish_places,omitempty"`
	Publishers         []string          `json:"publishers,omitempty"`
	Contributions      []string          `json:"contributions,omitempty"`
	DeweyDecimalClass  []string          `json:"dewey_decimal_class,omitempty"`
	Genres             []string          `json:"genres,omitempty"`
	LcClassifications  []string          `json:"lc_classifications,omitempty"`
	OtherTitles        []string          `json:"other_titles,omitempty"`
	Series             []string          `json:"series,omitempty"`
	SourceRecords      []string          `json:"source_records,omitempty"`
	Subjects           []string          `json:"subjects,omitempty"`
	WorkTitles         []string          `json:"work_titles,omitempty"`
	TableOfContents    []any             `json:"table_of_contents,omitempty"`
	Description        *TextBlock        `json:"description,omitempty"`
	FirstSentence      *TextBlock        `json:"first_sentence,omitempty"`
	Notes              *TextBlock        `json:"notes,omitempty"`
	Revision           int64             `json:"revision"`
	LatestRevision     int64             `json:"latest_revision,omitempty"`
	Created            *InternalDateTime `json:"created,omitempty"`
	LastModified       *InternalDateTime `json:"last_modified,omitempty"`
}

type EditionType struct {
	Key string `json:"key"`
}

type WorkReference struct {
	Key string `json:"key"`
}

type Language struct {
	Key string `json:"key"`
}

type EditionsResponse struct {
	Links struct {
		Self string `json:"self,omitempty"`
		Work string `json:"work,omitempty"`
		Next string `json:"next,omitempty"`
	} `json:"links"`
	Size    int       `json:"size,omitempty"`
	Entries []Edition `json:"entries"`
}

func (api *WorksAPI) Editions() (resp *EditionsResponse, err error) {
	_, err = api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		Get(path.Join(api.key, "editions") + ".json")

	if err != nil {
		return
	}
	return
}

type EditionAPI struct {
	openlibraryClient *Client
	key               string
}

func (c *Client) Edition(key string) *EditionAPI {
	api := EditionAPI{
		openlibraryClient: c,
		key:               key,
	}
	return &api
}

func (api *EditionAPI) Get() (resp *Edition, err error) {
	_, err = api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		Get(api.key + ".json")

	if err != nil {
		return
	}
	return
}
