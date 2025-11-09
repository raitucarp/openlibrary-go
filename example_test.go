package openlibrary_test

import (
	"fmt"

	"github.com/raitucarp/openlibrary-go"
)

// ExampleNewClient demonstrates creating a new OpenLibrary client.
func ExampleNewClient() {
	client := openlibrary.NewClient()
	_ = client
	fmt.Println("client ready")
	// Output: client ready
}

// ExampleAuthorSearch demonstrates searching authors.
func ExampleAuthorSearch() {
	client := openlibrary.NewClient()

	search, err := client.Authors().Query("Tolkien").Search()
	if err != nil || len(search.Docs) == 0 {
		fmt.Println("no results")
		return
	}

	first := search.Docs[0]
	fmt.Println(first.Name)
	// Output:
	// J.R.R. Tolkien
}

// ExampleAuthorWorks demonstrates fetching works by an author.
func ExampleAuthorWorks() {
	client := openlibrary.NewClient()

	search, err := client.Authors().Query("Brandon Sanderson").Search()
	if err != nil || len(search.Docs) == 0 {
		fmt.Println("no results")
		return
	}

	worksResp, err := search.Docs[0].Works().Fetch()
	if err != nil || worksResp == nil {
		fmt.Println("no works")
		return
	}

	fmt.Println(worksResp.Size > 0)
	// Output:
	// true
}

// ExampleCover demonstrates retrieving a cover from a work.
func ExampleCover() {
	client := openlibrary.NewClient()

	resp, err := client.Search().Query("Mistborn").
		Fields(openlibrary.TitleField, openlibrary.KeyField).
		Do()
	if err != nil || len(resp.Docs) == 0 {
		fmt.Println("no results")
		return
	}

	workClient := resp.ToClient()
	workAPI := workClient.Works(resp.Docs[0].Key)

	editions, err := workAPI.Editions()
	if err != nil || len(editions.Entries) == 0 {
		fmt.Println("no editions")
		return
	}

	edition := editions.Entries[0]
	book, err := client.Edition(edition.Key).Get()
	if err != nil {
		fmt.Println("no edition details")
		return
	}

	if len(book.ISBN10) == 0 {
		fmt.Println("no isbn")
		return
	}

	// We only demonstrate that API builds request correctly, no download.
	cover := client.Cover().ISBN(book.ISBN10[0]).Small()
	img, _, err := cover.Get()
	fmt.Println("cover request ready:", book.Title, len(img))
	// Output:
	// cover request ready: Mistborn
}

// ExampleSubjects demonstrates querying works by subject.
func ExampleSubjects() {
	client := openlibrary.NewClient()

	subjects, err := client.Subjects("love").Limit(5).Get()
	if err != nil || len(subjects.Works) == 0 {
		fmt.Println("no subject results")
		return
	}

	fmt.Println(len(subjects.Works) > 0)
	// Output:
	// true
}
