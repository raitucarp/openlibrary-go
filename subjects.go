package openlibrary

import (
	"fmt"
	"path"
	"strconv"
)

type SubjectsAPI struct {
	openlibraryClient *Client
	subject           string
	details           bool
	ebooks            bool
	publishedIn       []int
	limit             int
	offset            int
}

type SubjectsResponse struct {
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	SubjectType string `json:"subject_type,omitempty"`
	SolrQuery   string `json:"solr_query,omitempty"`
	WorkCount   int    `json:"work_count,omitempty"`
	Works       []struct {
		Key               string   `json:"key,omitempty"`
		Title             string   `json:"title,omitempty"`
		EditionCount      int      `json:"edition_count,omitempty"`
		CoverID           int      `json:"cover_id,omitempty"`
		CoverEditionKey   string   `json:"cover_edition_key,omitempty"`
		Subject           []string `json:"subject,omitempty"`
		IaCollection      []string `json:"ia_collection,omitempty"`
		Printdisabled     bool     `json:"printdisabled,omitempty"`
		LendingEdition    string   `json:"lending_edition,omitempty"`
		LendingIdentifier string   `json:"lending_identifier,omitempty"`
		Authors           []struct {
			Key  string `json:"key,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"authors,omitempty"`
		FirstPublishYear int    `json:"first_publish_year,omitempty"`
		Ia               string `json:"ia,omitempty"`
		PublicScan       bool   `json:"public_scan,omitempty"`
		HasFulltext      bool   `json:"has_fulltext,omitempty"`
		Availability     *struct {
			Status              string `json:"status,omitempty"`
			AvailableToBrowse   bool   `json:"available_to_browse,omitempty"`
			AvailableToBorrow   bool   `json:"available_to_borrow,omitempty"`
			AvailableToWaitlist bool   `json:"available_to_waitlist,omitempty"`
			IsPrintdisabled     bool   `json:"is_printdisabled,omitempty"`
			IsReadable          bool   `json:"is_readable,omitempty"`
			IsLendable          bool   `json:"is_lendable,omitempty"`
			IsPreviewable       bool   `json:"is_previewable,omitempty"`
			Identifier          string `json:"identifier,omitempty"`
			Isbn                any    `json:"isbn,omitempty"`
			Oclc                any    `json:"oclc,omitempty"`
			OpenlibraryWork     string `json:"openlibrary_work,omitempty"`
			OpenlibraryEdition  string `json:"openlibrary_edition,omitempty"`
			LastLoanDate        any    `json:"last_loan_date,omitempty"`
			NumWaitlist         any    `json:"num_waitlist,omitempty"`
			LastWaitlistDate    any    `json:"last_waitlist_date,omitempty"`
			IsRestricted        bool   `json:"is_restricted,omitempty"`
			IsBrowseable        bool   `json:"is_browseable,omitempty"`
			Src                 string `json:"__src__,omitempty"`
		} `json:"availability,omitempty"`
	} `json:"works,omitempty"`
	EbookCount int `json:"ebook_count,omitempty"`
	Subjects   []struct {
		Key   string `json:"key,omitempty"`
		Name  string `json:"name,omitempty"`
		Count int    `json:"count,omitempty"`
	} `json:"subjects,omitempty"`
	Places []struct {
		Key   string `json:"key,omitempty"`
		Name  string `json:"name,omitempty"`
		Count int    `json:"count,omitempty"`
	} `json:"places,omitempty"`
	People []struct {
		Key   string `json:"key,omitempty"`
		Name  string `json:"name,omitempty"`
		Count int    `json:"count,omitempty"`
	} `json:"people,omitempty"`
	Times []struct {
		Key   string `json:"key,omitempty"`
		Name  string `json:"name,omitempty"`
		Count int    `json:"count,omitempty"`
	} `json:"times,omitempty"`
	Authors []struct {
		Name  string `json:"name,omitempty"`
		Key   string `json:"key,omitempty"`
		Count int    `json:"count,omitempty"`
	} `json:"authors,omitempty"`
	Publishers []struct {
		Name  string `json:"name,omitempty"`
		Count int    `json:"count,omitempty"`
		Key   string `json:"key,omitempty"`
	} `json:"publishers,omitempty"`
	Languages []struct {
		Name  string `json:"name,omitempty"`
		Count int    `json:"count,omitempty"`
	} `json:"languages,omitempty"`
	PublishingHistory [][]int `json:"publishing_history,omitempty"`
}

func (c *Client) Subjects(subject string) *SubjectsAPI {
	api := SubjectsAPI{
		openlibraryClient: c,
		subject:           subject,
		limit:             10,
		offset:            0,
	}
	return &api
}

func (api *SubjectsAPI) WithDetails() *SubjectsAPI {
	api.details = true
	return api
}

func (api *SubjectsAPI) WithEbooks() *SubjectsAPI {
	api.ebooks = true
	return api
}

func (api *SubjectsAPI) PublishedIn(from int, to int) *SubjectsAPI {
	api.publishedIn = []int{from, to}
	return api
}

func (api *SubjectsAPI) Limit(limit int) *SubjectsAPI {
	api.limit = limit
	return api
}

func (api *SubjectsAPI) Offset(offset int) *SubjectsAPI {
	api.offset = offset
	return api
}

func (api *SubjectsAPI) Get() (resp *SubjectsResponse, err error) {
	endpoint := path.Join("/subjects", api.subject+".json")
	client := api.openlibraryClient.httpClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&resp).
		SetQueryParam("limit", strconv.Itoa(api.limit)).
		SetQueryParam("offset", strconv.Itoa(api.offset))

	if api.details {
		client.SetQueryParam("details", "true")
	}

	if len(api.publishedIn) == 2 {
		client.SetQueryParam("published_in", fmt.Sprintf("%d-%d", api.publishedIn[0], api.publishedIn[1]))
	}

	if api.ebooks {
		client.SetQueryParam("ebooks", "true")
	}

	_, err = client.Get(endpoint)

	// log.Println(endpoint, gg.Request.URL, string(gg.Bytes()))
	if err != nil {
		return
	}
	return
}
