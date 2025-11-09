package tests

import (
	"testing"

	"github.com/raitucarp/openlibrary-go"
)

func TestLoveSubjects(t *testing.T) {
	client := openlibrary.NewClient()
	loveSubject := client.Subjects("love")

	subjects, err := loveSubject.Get()
	if err != nil {
		t.Errorf("Something error %s", err)
	}

	if len(subjects.Works) <= 0 {
		t.Error("No works found with subject love")
	}
}

func TestWritingSubjectsDetails(t *testing.T) {
	client := openlibrary.NewClient()
	loveSubject := client.Subjects("writing").WithDetails()

	subjects, err := loveSubject.Get()

	if err != nil {
		t.Errorf("Something error %s", err)
	}

	if len(subjects.Works) <= 0 {
		t.Error("No works found with subject writing")
	}

	if len(subjects.Authors) <= 0 {
		t.Error("No authors found")
	}

}
