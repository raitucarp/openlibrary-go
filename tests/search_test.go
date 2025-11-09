package tests

import (
	"log"
	"testing"

	"github.com/raitucarp/openlibrary-go"
)

func TestSearch(t *testing.T) {
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

	ratings, err := workAPI.Ratings()
	bookshelves, err := workAPI.Bookshelves()

	log.Println("ratings", ratings)
	log.Println("bookshelves", bookshelves)

	if len(editions.Entries) <= 0 {
		t.Errorf("No editions")
		return
	}

	bookOne := editions.Entries[0]
	editionAPI := client.Edition(bookOne.Key)
	book, err := editionAPI.Get()
	if bookOne.Title != book.Title {
		t.Errorf("Title mismatch. Expected = %s, actual %s", bookOne.Title, book.Title)
	}

	bookByISBN, err := client.ISBN(book.ISBN13[0]).Get()

	if book.Title != bookByISBN.Title {
		t.Errorf("Title mismatch. Expected = %s, actual %s", book.Title, bookByISBN.Title)
	}
}
