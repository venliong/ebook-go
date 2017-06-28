package books

import (
	"fmt"
	"testing"
)

func TestGetEBook(t *testing.T) {
	fmt.Println(t.Name())
	book := GetEBook("./tmp/grdls.mobi")
	PrintEbook(book)
	book = GetEBook("./tmp/scm.epub")
	PrintEbook(book)
	book = GetEBook("./tmp/st1.azw3")
	PrintEbook(book)
}
