package mobi

import (
	"github.com/google/uuid"
	"github.com/jiangmitiao/books/utils"
	"io/ioutil"
	"os"
	"path"
)

type MobiBook struct {
	title            string
	author           string
	pubdate          string
	isbn             string
	publisher        string
	comments         string
	format           string
	uncompressedSize int64
	cover            []byte
	tags             []string

	FullPathFilename string
	Filename         string
	Ext              string
	TmpDir           string
	Enable           bool
}

func NewMobi(fullPathFilename string) *MobiBook {
	book := &MobiBook{
		FullPathFilename: fullPathFilename,
		Filename:         path.Base(fullPathFilename),
		Ext:              path.Ext(fullPathFilename),
		TmpDir:           path.Join(os.TempDir(), uuid.New().String()),
	}
	book.tags = make([]string, 0)
	return book
}

func (book *MobiBook) Parse() error {
	if b, err := ioutil.ReadFile(book.FullPathFilename); err == nil {
		meta := utils.NewMobiMeta()
		if err := meta.Parse(b); err == nil {
			book.title = meta.GetFullName()
			for _, value := range meta.MOBIHeader.EXTHHeader.EXTHRecords {
				switch value.Type() {
				case "publisher":
					book.publisher = value.Data()
					break
				case "description":
					book.comments = value.Data()
					break
				case "author":
					book.author = value.Data()
					break
				case "title":
					book.title = value.Data()
					break
				case "publishdate":
					book.pubdate = value.Data()
					break
				case "isbn":
					book.isbn = value.Data()
					break
				default:
					break
				}
			}
			book.format = "MOBI"
			if fileInfo, err := os.Stat(book.FullPathFilename); err == nil {
				book.uncompressedSize = fileInfo.Size()
			}
			book.cover = meta.GetCover()
			return nil
		} else {
			return err
		}
	} else {
		return err
	}

}

func (book *MobiBook) FullPath() string {
	return book.FullPathFilename
}

func (book *MobiBook) Title() string {
	return book.title
}
func (book *MobiBook) Author() string {
	return book.author
}
func (book *MobiBook) Pubdate() string {
	return book.pubdate
}
func (book *MobiBook) Isbn() string {
	return book.isbn
}
func (book *MobiBook) Comments() string {
	return book.comments
}
func (book *MobiBook) Format() string {
	return book.format
}
func (book *MobiBook) UncompressedSize() int64 {
	return book.uncompressedSize
}
func (book *MobiBook) Cover() ([]byte, bool) {
	if len(book.cover) != 0 {
		return book.cover, true
	}
	return book.cover, false
}
func (book *MobiBook) Publisher() string {
	return book.publisher
}
func (book *MobiBook) Tags() []string {
	return book.tags
}

func (book *MobiBook) Close() {
}
