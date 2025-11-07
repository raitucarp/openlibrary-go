package tests

import (
	"log"
	"os"
	"slices"
	"testing"

	"github.com/raitucarp/openlibrary-go"
)

func TestCover(t *testing.T) {
	client := openlibrary.NewClient()
	myQuery := "Mistborn"

	response, err := client.Search().Query(myQuery).
		Fields(openlibrary.TitleField, openlibrary.KeyField).
		Do()

	if err != nil {
		t.Errorf("Something error %s", err)
		return
	}

	if response.Query != myQuery {
		t.Errorf("Query is not equal, actual = %s, expected = %s", response.Query, myQuery)
		return
	}

	if len(response.Docs) <= 0 {
		return
	}

	client2 := response.ToClient()
	firstWorksKey := response.Docs[0].Key

	title := response.Docs[0].Title
	workAPI := client2.Works(firstWorksKey)
	works, err := workAPI.Get()
	if err != nil {
		t.Errorf("Something error %s", err)
	}

	if title != works.Title {
		t.Errorf("Title not equal. Expected = %s, actual = %s", title, works.Title)
		return
	}

	editions, err := workAPI.Editions()
	if err != nil {
		t.Errorf("Something error %s", err)
	}

	if len(editions.Entries) <= 0 {
		t.Errorf("No editions")
		return
	}

	withCoverIndex := slices.IndexFunc(editions.Entries, func(entry openlibrary.Edition) bool { return len(entry.Covers) > 0 })

	if withCoverIndex == -1 {
		t.Errorf("No books with cover")
	}

	bookOne := editions.Entries[withCoverIndex]
	editionAPI := client.Edition(bookOne.Key)
	book, err := editionAPI.Get()
	if bookOne.Title != book.Title {
		t.Errorf("Title mismatch. Expected = %s, actual %s", bookOne.Title, book.Title)
	}

	coverAPI := client.Cover()

	if len(book.ISBN10) > 0 {
		coverAPI.ISBN(book.ISBN10[0]).Medium()
	}

	if len(book.ISBN13) > 0 {
		coverAPI.ISBN(book.ISBN13[0]).Medium()
	}

	imgBytes, _, err := coverAPI.Get()

	if err != nil && len(book.LCCN) > 0 && len(imgBytes) <= 0 {
		coverAPI.LCCN(book.LCCN[0]).Small()
		imgBytes, _, err = coverAPI.Get()
	}

	if err != nil && len(book.OCLCNumbers) > 0 && len(imgBytes) <= 0 {
		coverAPI.OCLC(book.OCLCNumbers[0]).Small()
		imgBytes, _, err = coverAPI.Get()
	}

	if err != nil {
		t.Errorf("Something error %s", err)
		return
	}

	filePath := "cover.jpg"
	permissions := os.FileMode(0644)

	// Write the bytes to the file
	err = os.WriteFile(filePath, imgBytes, permissions)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}
