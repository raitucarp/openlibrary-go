package tests

import (
	"testing"

	"github.com/raitucarp/openlibrary-go"
)

func TestRead(t *testing.T) {
	client := openlibrary.NewClient()
	readAPI := client.Read()

	byISBN := readAPI.ISBN("0596156715")
	resp, err := byISBN.Get()
	if err != nil {
		t.Errorf("Something error %s", err)
	}

	if len(resp.Items) <= 0 {
		t.Error("No items")
	}

	if len(resp.Records) <= 0 {
		t.Error("No records")
	}

}

func TestReadBySearch(t *testing.T) {
	client := openlibrary.NewClient()
	searchClient := client.Search()

	searchQ := searchClient.Query("The Art of")

	result, err := searchQ.Do()
	if err != nil {
		t.Error("No items")
	}

	if result.NumFound <= 0 {
		t.Error("Numfound 0")
	}

	firstDoc := result.Docs[0]

	readBook, err := firstDoc.ReadByISBN()

	if readBook == nil {
		readBook, err = firstDoc.ReadByLCCN()
	}

	if readBook == nil {
		readBook, err = firstDoc.ReadByOCLC()
	}

	if readBook == nil {
		readBook, err = firstDoc.Read()
	}

	if readBook != nil {
		_, err := readBook.Get()
		if err != nil {
			t.Error(err)
		}
	}
}
