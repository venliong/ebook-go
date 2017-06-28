package epub

import (
	"fmt"
	"testing"
)

func TestEpubBook_Parse(t *testing.T) {
	fmt.Println(t.Name())
	book := NewEpub("../tmp/scm.epub")
	book.Parse()
	defer book.Close()
	fmt.Println(book.Parse())
	fmt.Printf("%+v \n", book)
}
