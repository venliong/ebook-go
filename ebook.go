package books

import (
	"fmt"
	"github.com/jiangmitiao/ebook-go/epub"
	"github.com/jiangmitiao/ebook-go/mobi"
	"strconv"
)

type EbookFile interface {
	Parse() error
}

type Ebook interface {
	FullPath() string
	Title() string
	Author() string
	Pubdate() string
	Isbn() string
	Comments() string
	Format() string
	UncompressedSize() int64
	Cover() ([]byte, bool)
	Publisher() string
	Tags() []string
}

func GetEBook(fullpathname string) Ebook {
	book := epub.NewEpub(fullpathname)
	if err := book.Parse(); err == nil {
		return book
	} else {
		//fmt.Println(err)
	}

	book1 := mobi.NewMobi(fullpathname)
	if err := book1.Parse(); err == nil {
		return book1
	} else {
		//fmt.Println(err)
	}
	return nil
}

func PrintEbook(v Ebook) {
	if v == nil {
		fmt.Println("nil")
		return
	}
	strs := ""
	strs += "Title:"
	strs += v.Title() + "\r\n"
	strs += "Author:"
	strs += v.Author() + "\r\n"
	strs += "Comments:"
	strs += v.Comments() + "\r\n"
	strs += "CoverSize:"
	tmp, _ := v.Cover()
	strs += strconv.FormatInt(int64(len(tmp)), 10) + "\r\n"
	fmt.Println(strs)
}
