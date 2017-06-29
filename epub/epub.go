package epub

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jiangmitiao/ebook-go/utils"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	defaultmimetype = "application/epub+zip"
)

type EpubBook struct {
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

	zip *utils.Zip
}

func NewEpub(fullPathFilename string) *EpubBook {
	book := &EpubBook{
		FullPathFilename: fullPathFilename,
		Filename:         path.Base(fullPathFilename),
		Ext:              path.Ext(fullPathFilename),
		TmpDir:           path.Join(os.TempDir(), uuid.New().String()),
	}
	book.tags = make([]string, 0)
	return book
}

func (book *EpubBook) Parse() error {
	book.zip = utils.NewZipWithTmpDir(book.FullPathFilename, book.TmpDir)
	if err := book.zip.Decompress(); err == nil {
		if mimetype, err := ioutil.ReadFile(path.Join(book.TmpDir, "mimetype")); err == nil {
			mimetypestr := strings.Trim(string(mimetype[:]), "\r\n")
			if strings.Compare(defaultmimetype, mimetypestr) == 0 {
				book.parse()
				return nil
			} else {
				book.Enable = false
				book.Close()
				return errors.New("not a epub " + mimetypestr)
			}
		} else {
			book.Enable = false
			book.Close()
			return err
		}
	} else {
		book.Enable = false
		book.Close()
		return err
	}
}

func (book *EpubBook) parse() {
	container := utils.NewXml(path.Join(book.TmpDir, "META-INF", "container.xml"))
	if pathname, err := container.GetAttr("//container/rootfiles/rootfile", "full-path"); err == nil {
		book.Enable = true
		content := utils.NewXml(path.Join(book.TmpDir, pathname))
		book.parseData()
		book.parseCover(content)
		book.parseBook(content)
		book.parseAuthor(content)
		book.parseLanguage(content)
		book.parsePublish(content)
		book.parseComments(content)
		book.parseTag(content)
	}
}

func (book *EpubBook) parseData() {
	book.format = "EPUB"
	if fileInfo, err := os.Stat(book.FullPathFilename); err == nil {
		book.uncompressedSize = fileInfo.Size()
	}
}

func (book *EpubBook) parseCover(content *utils.Xml) {
	if _, err := content.GetValue("//package/metadata/meta[@name='cover']"); err == nil {
		if metaid, _ := content.GetAttr("//package/metadata/meta[@name='cover']", "content"); metaid != "" {
			paths, _ := content.GetAttr("//package/manifest/item[@id='"+metaid+"']", "href")
			fullpath := path.Join(path.Dir(content.FullPathFilename), paths)
			book.cover, _ = ioutil.ReadFile(fullpath)
		}
	}
}

func (book *EpubBook) parseBook(content *utils.Xml) {
	book.title, _ = content.GetValue("//package/metadata/dc:title")
	book.pubdate, _ = content.GetValue("//package/metadata/dc:date")
	book.isbn, _ = content.GetValue("//package/metadata/dc:identifier[@scheme='ISBN']")
}
func (book *EpubBook) parseAuthor(content *utils.Xml) {
	book.author, _ = content.GetValue("//package/metadata/dc:creator")
}

func (book *EpubBook) parseLanguage(content *utils.Xml) {
	//book.Language.Lang, _ = content.GetValue("//package/metadata/dc:language")
}
func (book *EpubBook) parsePublish(content *utils.Xml) {
	book.publisher, _ = content.GetValue("//package/metadata/dc:publisher")
}

func (book *EpubBook) parseComments(content *utils.Xml) {
	book.comments, _ = content.GetValue("//package/metadata/dc:description")
}

func (book *EpubBook) parseTag(content *utils.Xml) {
	book.tags, _ = content.GetValues("//package/metadata/dc:subject")
}

func (book *EpubBook) FullPath() string {
	return book.FullPathFilename
}

func (book *EpubBook) Title() string {
	return book.title
}
func (book *EpubBook) Author() string {
	return book.author
}
func (book *EpubBook) Pubdate() string {
	return book.pubdate
}
func (book *EpubBook) Isbn() string {
	return book.isbn
}
func (book *EpubBook) Comments() string {
	return book.comments
}
func (book *EpubBook) Format() string {
	return book.format
}
func (book *EpubBook) UncompressedSize() int64 {
	return book.uncompressedSize
}
func (book *EpubBook) Cover() ([]byte, bool) {
	if len(book.cover) != 0 {
		return book.cover, true
	}
	return book.cover, false
}
func (book *EpubBook) Publisher() string {
	return book.publisher
}
func (book *EpubBook) Tags() []string {
	return book.tags
}

func (book *EpubBook) Close() {
	book.zip.Close()
}
