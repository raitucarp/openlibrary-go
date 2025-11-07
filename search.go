package openlibrary

import (
	"net/url"
	"strings"
)

type SearchField string

const (
	KeyField                   SearchField = "key"
	RedirectsField             SearchField = "redirects"
	TitleField                 SearchField = "title"
	SubtitleField              SearchField = "subtitle"
	AlternativeTitleField      SearchField = "alternative_title"
	AlternativeSubtitleField   SearchField = "alternative_subtitle"
	CoverIField                SearchField = "cover_i"
	EbookAccessField           SearchField = "ebook_access"
	EditionCountField          SearchField = "edition_count"
	EditionKeyField            SearchField = "edition_key"
	FormatField                SearchField = "format"
	ByStatementField           SearchField = "by_statement"
	PublishDateField           SearchField = "publish_date"
	LccnField                  SearchField = "lccn"
	IAField                    SearchField = "ia"
	OCLCField                  SearchField = "oclc"
	ISBNField                  SearchField = "isbn"
	ContributorField           SearchField = "contributor"
	PublishPlaceField          SearchField = "publish_place"
	PublisherField             SearchField = "publisher"
	FirstSentenceField         SearchField = "first_sentence"
	AuthorKeyField             SearchField = "author_key"
	AuthorNameField            SearchField = "author_name"
	AuthorAlternativeNameField SearchField = "author_alternative_name"
	SubjectField               SearchField = "subject"
	PersonField                SearchField = "person"
	PlaceField                 SearchField = "place"
	TimeField                  SearchField = "time"
	HasFulltextField           SearchField = "has_fulltext"
	TitleSuggestField          SearchField = "title_suggest"
	PublishYearField           SearchField = "publish_year"
	LanguageField              SearchField = "language"
	NumberOfPagesMedianField   SearchField = "number_of_pages_median"
	IACountField               SearchField = "ia_count"
	PublisherFacetField        SearchField = "publisher_facet"
	AuthorFacetField           SearchField = "author_facet"
	FirstPublishYearField      SearchField = "first_publish_year"
	RatingsCountField          SearchField = "ratings_count"
	ReadinglogCountField       SearchField = "readinglog_count"
	WantToReadCountField       SearchField = "want_to_read_count"
	CurrentlyReadingCountField SearchField = "currently_reading_count"
	AlreadyReadCountField      SearchField = "already_read_count"
	SubjectKeyField            SearchField = "subject_key"
	PersonKeyField             SearchField = "person_key"
	PlaceKeyField              SearchField = "place_key"
	TimeKeyField               SearchField = "time_key"
	LLCField                   SearchField = "lcc"
	DDCField                   SearchField = "ddc"
	LCCSortField               SearchField = "lcc_sort"
	DDCSortField               SearchField = "ddc_sort"
)

type EbookAccess string

const (
	NoEbookAccess       EbookAccess = "no_ebook"
	UnclassifiedAccess  EbookAccess = "unclassified"
	PrintDisabledAccess EbookAccess = "printdisabled"
	BorrowableAccess    EbookAccess = "borrowable"
	PublicAccessAccess  EbookAccess = "public"
)

func (a EbookAccess) String() string {
	return string(a)
}

type SearchAPI struct {
	openlibraryClient *Client
	fields            []SearchField
	queryKey          string
}

type SearchResult struct {
	Key                   string      `json:"key"`
	Redirects             []string    `json:"redirects,omitempty"`
	Title                 string      `json:"title"`
	Subtitle              string      `json:"subtitle,omitempty"`
	AlternativeTitle      []string    `json:"alternative_title,omitempty"`
	AlternativeSubtitle   []string    `json:"alternative_subtitle,omitempty"`
	CoverI                int         `json:"cover_i,omitempty"`
	EbookAccess           EbookAccess `json:"ebook_access,omitempty"`
	EditionCount          int         `json:"edition_count,omitempty"`
	EditionKey            []string    `json:"edition_key,omitempty"`
	Format                []string    `json:"format,omitempty"`
	ByStatement           []string    `json:"by_statement,omitempty"`
	PublishDate           []string    `json:"publish_date,omitempty"`
	LCCN                  []string    `json:"lccn,omitempty"`
	IA                    []string    `json:"ia,omitempty"`
	Oclc                  []string    `json:"oclc,omitempty"`
	ISBN                  []string    `json:"isbn,omitempty"`
	Contributor           []string    `json:"contributor,omitempty"`
	PublishPlace          []string    `json:"publish_place,omitempty"`
	Publisher             []string    `json:"publisher,omitempty"`
	FirstSentence         []string    `json:"first_sentence,omitempty"`
	AuthorKey             []string    `json:"author_key,omitempty"`
	AuthorName            []string    `json:"author_name,omitempty"`
	AuthorAlternativeName []string    `json:"author_alternative_name,omitempty"`
	Subject               []string    `json:"subject,omitempty"`
	Person                []string    `json:"person,omitempty"`
	Place                 []string    `json:"place,omitempty"`
	Time                  []string    `json:"time"`
	HasFulltext           bool        `json:"has_fulltext,omitempty"`
	TitleSuggest          string      `json:"title_suggest,omitempty"`
	PublishYear           []int       `json:"publish_year,omitempty"`
	Language              []string    `json:"language,omitempty"`
	NumberOfPagesMedian   int         `json:"number_of_pages_median,omitempty"`
	IaCount               int         `json:"ia_count,omitempty"`
	PublisherFacet        []string    `json:"publisher_facet,omitempty"`
	AuthorFacet           []string    `json:"author_facet,omitempty"`
	FirstPublishYear      int         `json:"first_publish_year,omitempty"`
	RatingsCount          int         `json:"ratings_count,omitempty"`
	ReadinglogCount       int         `json:"readinglog_count,omitempty"`
	WantToReadCount       int         `json:"want_to_read_count,omitempty"`
	CurrentlyReadingCount int         `json:"currently_reading_count,omitempty"`
	AlreadyReadCount      int         `json:"already_read_count,omitempty"`
	SubjectKey            []string    `json:"subject_key,omitempty"`
	PersonKey             []string    `json:"person_key,omitempty"`
	PlaceKey              []string    `json:"place_key,omitempty"`
	TimeKey               []string    `json:"time_key,omitempty"`
	LCC                   []string    `json:"lcc,omitempty"`
	DDC                   []string    `json:"ddc,omitempty"`
	LCCSort               string      `json:"lcc_sort,omitempty"`
	DDCSort               string      `json:"ddc_sort,omitempty"`
}

type SearchResponse struct {
	NumFound      int            `json:"numFound,omitempty"`
	Start         int            `json:"start,omitempty"`
	NumFoundExact bool           `json:"numFoundExact,omitempty"`
	NumFound_     int            `json:"num_found,omitempty"`
	Query         string         `json:"q,omitempty"`
	DocUrl        string         `json:"documentation_url,omitempty"`
	Offset        int            `json:"offset,omitempty"`
	Docs          []SearchResult `json:"docs,omitempty"`

	openlibraryClient *Client
}

func (sr *SearchResponse) ToClient() *Client {
	return sr.openlibraryClient
}

func (c *Client) Search() *SearchAPI {
	s := SearchAPI{
		openlibraryClient: c,
	}
	return &s
}

func (api *SearchAPI) Query(q string) *SearchAPI {
	api.queryKey = q
	return api
}

func (api *SearchAPI) Fields(fields ...SearchField) *SearchAPI {
	api.fields = append(api.fields, fields...)
	return api
}

func (api *SearchAPI) Author(author string) *SearchAPI {
	return api
}

func (api *SearchAPI) Authors(author string) *SearchAPI {
	return api
}

func (api *SearchAPI) Do() (resp *SearchResponse, err error) {
	fieldsCandidate := []string{}
	for _, field := range api.fields {
		fieldsCandidate = append(fieldsCandidate, string(field))
	}

	q := url.Values{}

	if api.queryKey != "" {
		q.Add("q", api.queryKey)
	}

	if len(fieldsCandidate) > 0 {
		q.Add("fields", strings.Join(fieldsCandidate, ","))
	}

	_, err = api.openlibraryClient.httpClient.R().
		SetQueryString(q.Encode()).
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		Get("/search")

	if err != nil {
		return
	}

	resp.openlibraryClient = api.openlibraryClient

	return resp, err
}
