package openlibrary

import (
	"path"
)

type Identifiers struct {
	Abaa                                        []string `json:"abaa,omitempty"`
	DominicanInstituteForOrientalStudiesLibrary []string `json:"dominican_institute_for_oriental_studies_library,omitempty"`
	AlibrisId                                   []string `json:"alibris_id,omitempty"`
	Amazon                                      []string `json:"amazon,omitempty"`
	AbwaBibliographicNumber                     []string `json:"abwa_bibliographic_number,omitempty"`
	BetterWorldBooks                            []string `json:"better_world_books,omitempty"`
	DepositoLegal                               []string `json:"depósito_legal,omitempty"`
	BibliothequeNationaleDeFrance               []string `json:"bibliothèque_nationale_de_france,omitempty"`
	Bibsys                                      []string `json:"bibsys,omitempty"`
	Bhl                                         []string `json:"bhl,omitempty"`
	Bookbrainz                                  []string `json:"bookbrainz,omitempty"`
	Bcid                                        []string `json:"bcid,omitempty"`
	BooklockerCom                               []string `json:"booklocker.com,omitempty"`
	Bookmooch                                   []string `json:"bookmooch,omitempty"`
	Bookwire                                    []string `json:"bookwire,omitempty"`
	Booksforyou                                 []string `json:"booksforyou,omitempty"`
	BostonPublicLibrary                         []string `json:"boston_public_library,omitempty"`
	BritishLibrary                              []string `json:"british_library,omitempty"`
	CornellUniversityOnlineLibrary              []string `json:"cornell_university_online_library,omitempty"`
	CornellUniversityLibrary                    []string `json:"cornell_university_library,omitempty"`
	CanadianNationalLibraryArchive              []string `json:"canadian_national_library_archive,omitempty"`
	Choosebooks                                 []string `json:"choosebooks,omitempty"`
	Dnb                                         []string `json:"dnb,omitempty"`
	DigitalLibraryPomerania                     []string `json:"digital_library_pomerania,omitempty"`
	Doi                                         []string `json:"doi,omitempty"`
	Discovereads                                []string `json:"discovereads,omitempty"`
	Freebase                                    []string `json:"freebase,omitempty"`
	Folio                                       []string `json:"folio,omitempty"`
	Goodreads                                   []string `json:"goodreads,omitempty"`
	Google                                      []string `json:"google,omitempty"`
	GrandComicsDatabase                         []string `json:"grand_comics_database,omitempty"`
	HathiTrust                                  []string `json:"hathi_trust,omitempty"`
	Harvard                                     []string `json:"harvard,omitempty"`
	Ilmiolibro                                  []string `json:"ilmiolibro,omitempty"`
	Inducks                                     []string `json:"inducks,omitempty"`
	Ocaid                                       []string `json:"ocaid,omitempty"`
	Isfdb                                       []string `json:"isfdb,omitempty"`
	Etsc                                        []string `json:"etsc,omitempty"`
	Isbn10                                      []string `json:"isbn_10,omitempty"`
	Isbn13                                      []string `json:"isbn_13,omitempty"`
	Issn                                        []string `json:"issn,omitempty"`
	Istc                                        []string `json:"istc,omitempty"`
	Lccn                                        []string `json:"lccn,omitempty"`
	Librarything                                []string `json:"librarything,omitempty"`
	Lulu                                        []string `json:"lulu,omitempty"`
	Magcloud                                    []string `json:"magcloud,omitempty"`
	Musicbrainz                                 []string `json:"musicbrainz,omitempty"`
	Nla                                         []string `json:"nla,omitempty"`
	Nbuv                                        []string `json:"nbuv,omitempty"`
	Libris                                      []string `json:"libris,omitempty"`
	OclcNumbers                                 []string `json:"oclc_numbers,omitempty"`
	Overdrive                                   []string `json:"overdrive,omitempty"`
	PaperbackSwap                               []string `json:"paperback_swap,omitempty"`
	ProjectGutenberg                            []string `json:"project_gutenberg,omitempty"`
	ProjectRuneberg                             []string `json:"project_runeberg,omitempty"`
	Scribd                                      []string `json:"scribd,omitempty"`
	OpacSbn                                     []string `json:"opac_sbn,omitempty"`
	Shelfari                                    []string `json:"shelfari,omitempty"`
	SmashwordsBookDownload                      []string `json:"smashwords_book_download,omitempty"`
	StandardEbooks                              []string `json:"standard_ebooks,omitempty"`
	Storygraph                                  []string `json:"storygraph,omitempty"`
	Ulrls                                       []string `json:"ulrls,omitempty"`
	WWNorton                                    []string `json:"w._w._norton,omitempty"`
	ZdbId                                       []string `json:"zdb-id,omitempty"`
	Fennica                                     []string `json:"fennica,omitempty"`
	BayerischeStaatsbibliothek                  []string `json:"bayerische_staatsbibliothek,omitempty"`
	AbebooksDe                                  []string `json:"abebooks.de,omitempty"`
	DcBooks                                     []string `json:"dc_books,omitempty"`
	Publishamerica                              []string `json:"publishamerica,omitempty"`
	BritishNationalBibliography                 []string `json:"british_national_bibliography,omitempty"`
	Wikidata                                    []string `json:"wikidata,omitempty"`
	Librivox                                    []string `json:"librivox,omitempty"`
	OpenAlex                                    []string `json:"open_alex,omitempty"`
	Openstax                                    []string `json:"openstax,omitempty"`
	OpenTextbookLibrary                         []string `json:"open_textbook_library,omitempty"`
	Wikisource                                  []string `json:"wikisource,omitempty"`
	Yakaboo                                     []string `json:"yakaboo,omitempty"`
	Infosoup                                    []string `json:"infosoup,omitempty"`
	UrnNbn                                      []string `json:"urn_nbn,omitempty"`
}

type Edition struct {
	Key                string            `json:"key"`
	Title              string            `json:"title"`
	Subtitle           *string           `json:"subtitle,omitempty"`
	Type               EditionType       `json:"type"`
	Authors            []AuthorWithKey   `json:"authors,omitempty"`
	Works              []WorkReference   `json:"works"`
	Identifiers        *Identifiers      `json:"identifiers,omitempty"`
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
	PublishCountry     string            `json:"publish_country,omitempty"`
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
