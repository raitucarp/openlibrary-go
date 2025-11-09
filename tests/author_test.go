package tests

import (
	"log"
	"os"
	"testing"

	"github.com/raitucarp/openlibrary-go"
)

func TestSearchAuthor(t *testing.T) {
	client := openlibrary.NewClient()
	authors := client.Authors()

	authorQuery := authors.Query("Tolkien")
	searchResult, err := authorQuery.Search()
	if err != nil {
		t.Errorf("Something error %s", err)
	}

	firstAuthor := searchResult.Docs[0]

	imgBytes, _, err := firstAuthor.Photo().Get()

	if err != nil {
		t.Errorf("Something error %s", err)
		return
	}

	filePath := "author_tolkien.jpg"
	permissions := os.FileMode(0644)

	// Write the bytes to the file
	err = os.WriteFile(filePath, imgBytes, permissions)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

func TestAuthorPhoto(t *testing.T) {
	client := openlibrary.NewClient()
	authors := client.Authors()

	authorQuery := authors.Query("Brandon Sanderson")
	searchResult, err := authorQuery.Search()
	if err != nil {
		t.Errorf("Something error %s", err)
	}

	firstAuthor := searchResult.Docs[0]
	imgBytes, _, err := authors.Photo().OLID(firstAuthor.Key).Get()

	if err != nil {
		t.Errorf("Something error %s", err)
		return
	}

	filePath := "author_brandon_sanderson.jpg"
	permissions := os.FileMode(0644)

	// Write the bytes to the file
	err = os.WriteFile(filePath, imgBytes, permissions)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

func TestAuthorWorks(t *testing.T) {
	client := openlibrary.NewClient()
	authors := client.Authors()

	authorQuery := authors.Query("Robert Jordan")
	searchResult, err := authorQuery.Search()
	if err != nil {
		t.Errorf("Something error %s", err)
	}

	firstAuthor := searchResult.Docs[0]
	workOfAuthor := firstAuthor.Works()

	works, err := workOfAuthor.Fetch()
	if err != nil {
		t.Errorf("Something error %s", err)
	}

	log.Println(works)
}
