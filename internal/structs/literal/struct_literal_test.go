package literal

import "testing"

func TestBookLiteral(t *testing.T) {
	book1 := Book{
		Title:     "The Go Programming Language",
		Author:    "Alan A. A. Donovan",
		Year:      2015,
		PageCount: 380,
	}

	book2 := struct {
		Title  string
		Author string
		Pages  int
	}{
		Title:  "Clean Code",
		Author: "Robert C. Martin",
		Pages:  464,
	}

	if book1.Title != "The Go Programming Language" {
		t.Errorf("unexpected title: %s", book1.Title)
	}
	if book2.Author != "Robert C. Martin" {
		t.Errorf("unexpected author: %s", book2.Author)
	}
	if book2.Pages != 464 {
		t.Errorf("unexpected pages count: %d", book2.Pages)
	}
}
