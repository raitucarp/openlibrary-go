package openlibrary

import (
	"errors"
	"net/http"
	"path"
)

type AuthorPictureKind string

const (
	AuthorByID   AuthorPictureKind = "id"
	AuthorByOLID AuthorPictureKind = "olid"
)

type AuthorSize string

const (
	AuthorPictureSmall  AuthorSize = "S"
	AuthorPictureMedium AuthorSize = "M"
	AuthorPictureLarge  AuthorSize = "L"
)

type AuthorsAPI struct {
	openlibraryClient *Client
}

type Author struct {
	Key            string            `json:"key"`
	Name           string            `json:"name"`
	Type           string            `json:"type"`
	AlternateNames []string          `json:"alternate_names,omitempty"`
	Bio            *TextBlock        `json:"bio,omitempty"`
	BirthDate      *string           `json:"birth_date,omitempty"`
	DeathDate      *string           `json:"death_date,omitempty"`
	Date           *string           `json:"date,omitempty"`
	EntityType     *string           `json:"entity_type,omitempty"`
	FullerName     *string           `json:"fuller_name,omitempty"`
	PersonalName   *string           `json:"personal_name,omitempty"`
	Title          *string           `json:"title,omitempty"`
	Photos         []int64           `json:"photos,omitempty"`
	Links          []Link            `json:"links,omitempty"`
	RemoteIDs      *RemoteIDs        `json:"remote_ids,omitempty"`
	Revision       int64             `json:"revision"`
	LatestRevision *int64            `json:"latest_revision,omitempty"`
	Created        *InternalDateTime `json:"created,omitempty"`
	LastModified   InternalDateTime  `json:"last_modified"`
}

type AuthorType struct {
	Key string `json:"key"`
}

type RemoteIDs struct {
	Wikidata *string `json:"wikidata,omitempty"`
	Viaf     *string `json:"viaf,omitempty"`
}

type AuthorSearch struct {
	openlibraryClient *Client
	query             string
}

type AuthorSearchResult struct {
	Key                   string   `json:"key,omitempty"`
	Name                  string   `json:"name,omitempty"`
	AlternateNames        []string `json:"alternate_names,omitempty"`
	BirthDate             string   `json:"birth_date,omitempty"`
	DeathDate             string   `json:"death_date,omitempty"`
	Date                  string   `json:"date,omitempty"`
	WorkCount             int      `json:"work_count,omitempty"`
	TopWork               string   `json:"top_work,omitempty"`
	TopSubjects           []string `json:"top_subjects,omitempty"`
	Type                  string   `json:"type,omitempty"`
	RatingsAverage        float64  `json:"ratings_average,omitempty"`
	RatingsSortable       float64  `json:"ratings_sortable,omitempty"`
	RatingsCount          int      `json:"ratings_count,omitempty"`
	RatingsCount1         int      `json:"ratings_count_1,omitempty"`
	RatingsCount2         int      `json:"ratings_count_2,omitempty"`
	RatingsCount3         int      `json:"ratings_count_3,omitempty"`
	RatingsCount4         int      `json:"ratings_count_4,omitempty"`
	RatingsCount5         int      `json:"ratings_count_5,omitempty"`
	WantToReadCount       int      `json:"want_to_read_count,omitempty"`
	AlreadyReadCount      int      `json:"already_read_count,omitempty"`
	CurrentlyReadingCount int      `json:"currently_reading_count,omitempty"`
	ReadinglogCount       int      `json:"readinglog_count,omitempty"`
	Version               int64    `json:"_version_,omitempty"`
}

type AuthorSearchResponse struct {
	NumFound      int                  `json:"numFound,omitempty"`
	Start         int                  `json:"start,omitempty"`
	NumFoundExact bool                 `json:"numFoundExact,omitempty"`
	Docs          []AuthorSearchResult `json:"docs,omitempty"`
}

func (c *Client) Authors() *AuthorsAPI {
	api := &AuthorsAPI{
		openlibraryClient: c,
	}

	return api
}

func (api *AuthorsAPI) Query(author string) *AuthorSearch {
	return &AuthorSearch{
		openlibraryClient: api.openlibraryClient,
		query:             author,
	}
}

func (api *AuthorSearch) Search() (resp *AuthorSearchResponse, err error) {
	_, err = api.openlibraryClient.httpClient.
		SetBaseURL(baseURL).R().
		SetHeader("Accept", "application/json").
		SetQueryParam("q", api.query).
		SetResult(&resp).
		Get("/search/authors")

	if err != nil {
		return
	}

	return
}

type AuthorAPI struct {
	openlibraryClient *Client
	identifier        string
	data              *Author
}

func (api *AuthorsAPI) ByIdentifier(identifier string) *AuthorAPI {
	return &AuthorAPI{
		openlibraryClient: api.openlibraryClient,
		identifier:        identifier,
	}
}

func (api *AuthorAPI) Fetch() (err error) {
	_, err = api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&api.data).
		Get(api.identifier + ".json")

	if err != nil {
		return
	}

	return
}

func (api *AuthorAPI) Data() (author *Author) {
	return api.data
}

func (api *AuthorAPI) Photo() (photoAPI *AuthorPhotoAPI) {
	if api.data == nil {
		return
	}
	_, value := path.Split(api.data.Key)

	return &AuthorPhotoAPI{
		openlibraryClient: api.openlibraryClient,
		kind:              AuthorByOLID,
		size:              AuthorPictureMedium,
		value:             value,
	}
}

type AuthorPhotoAPI struct {
	openlibraryClient *Client
	value             string
	kind              AuthorPictureKind
	size              AuthorSize
}

func (api *Author) Photo() (photoAPI *AuthorPhotoAPI) {
	_, value := path.Split(api.Key)

	return &AuthorPhotoAPI{
		openlibraryClient: NewClient(),
		kind:              AuthorByOLID,
		size:              AuthorPictureMedium,
		value:             value,
	}
}

func (api *AuthorSearchResult) Photo() (photoAPI *AuthorPhotoAPI) {
	_, value := path.Split(api.Key)

	return &AuthorPhotoAPI{
		openlibraryClient: NewClient(),
		kind:              AuthorByOLID,
		size:              AuthorPictureMedium,
		value:             value,
	}
}

func (api *AuthorsAPI) Photo() *AuthorPhotoAPI {
	return &AuthorPhotoAPI{
		openlibraryClient: api.openlibraryClient,
		kind:              AuthorByID,
		size:              AuthorPictureMedium,
	}
}

func (api *AuthorPhotoAPI) ID(id string) *AuthorPhotoAPI {
	api.value = id
	api.kind = AuthorByID
	return api
}

func (api *AuthorPhotoAPI) OLID(olid string) *AuthorPhotoAPI {
	api.value = olid
	api.kind = AuthorByOLID
	return api
}

func (api *AuthorPhotoAPI) Small() *AuthorPhotoAPI {
	api.size = AuthorPictureSmall
	return api
}

func (api *AuthorPhotoAPI) Medium() *AuthorPhotoAPI {
	api.size = AuthorPictureMedium
	return api
}

func (api *AuthorPhotoAPI) Large() *AuthorPhotoAPI {
	api.size = AuthorPictureLarge
	return api
}

func (api *AuthorPhotoAPI) Get() (imgBytes []byte, mimeType string, err error) {
	if api.value == "" {
		return imgBytes, mimeType, errors.New("No value")
	}

	endpoint := path.Join("/a", string(api.kind), api.value+"-"+string(api.size)) + ".jpg?default=false"

	res, err := api.openlibraryClient.httpClient.SetBaseURL(coverBaseURL).R().Get(endpoint)
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

type AuthorWorksAPI struct {
	openlibraryClient *Client
	authorKey         string
	limit             int
	offset            int
}

func (api *Author) Works() (worksAPI *AuthorWorksAPI) {
	_, authorKey := path.Split(api.Key)

	return &AuthorWorksAPI{
		openlibraryClient: NewClient(),
		authorKey:         authorKey,
		limit:             10,
		offset:            0,
	}
}

func (api *AuthorSearchResult) Works() (worksAPI *AuthorWorksAPI) {
	_, authorKey := path.Split(api.Key)
	return &AuthorWorksAPI{
		openlibraryClient: NewClient(),
		authorKey:         authorKey,
		limit:             10,
		offset:            0,
	}
}

func (api *AuthorWorksAPI) Limit(limit int) *AuthorWorksAPI {
	api.limit = limit
	return api
}

func (api *AuthorWorksAPI) Offset(offset int) *AuthorWorksAPI {
	api.offset = offset
	return api
}

type AuthorWorksResponse struct {
	Links struct {
		Self string `json:"self,omitempty"`
		Work string `json:"work,omitempty"`
		Next string `json:"next,omitempty"`
	} `json:"links"`
	Size    int     `json:"size,omitempty"`
	Entries []Works `json:"entries"`
}

func (api *AuthorWorksAPI) Fetch() (resp *AuthorWorksResponse, err error) {
	_, err = api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		Get(path.Join("/authors", api.authorKey, "works.json"))

	if err != nil {
		return
	}

	return
}
