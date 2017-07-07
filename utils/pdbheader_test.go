package utils

import (
	"bytes"
	"fmt"
	"github.com/jiangmitiao/ebook-go/utils/lz77"
	_ "github.com/jiangmitiao/ebook-go/utils/lz77"
	"io"
	"io/ioutil"
	"testing"
	"unicode/utf8"
	"strconv"
)

func TestPDBHeader_Parse(t *testing.T) {
	fmt.Println(t.Name())
	header := NewPDBHeader()
	//b, _ := ioutil.ReadFile("../tmp/grdls.mobi")
	b, _ := ioutil.ReadFile("../tmp/zcdz.mobi")
	//b,_ :=ioutil.ReadFile("/home/gavin/Calibre 书库/Wang Ceng Qi/Shou Jie (3)/sj.mobi")
	//b, _ := ioutil.ReadFile("../tmp/st1.azw3")
	fmt.Println("file size : "+strconv.Itoa(len(b)))
	reader := bytes.NewReader(b)
	header.Parse(reader)

	var tmp = to32Byte(header.Name)
	fmt.Println("header name :"+string(tmp))
	fmt.Println("record nums ",strconv.Itoa(int(header.GetNumRecords())))
	fmt.Println("attributes ",header.GetAttributesMean(),header.GetAttributes())
	fmt.Println("file version ",header.Version)
	fmt.Println("CreationDate",header.CreationDate)
	fmt.Println("ModificationDate",header.ModificationDate)


	for _, value := range header.RecordInfos {
		fmt.Print(value.GetUniqueID())
	}
	fmt.Println("header size ", header.Size())
	fmt.Println("mobi header size ", header.GetMobiHeaderSize())

	start, offset := header.GetRecord(1)
	fmt.Println("start ", start, "offset", offset)
	tmp = make([]byte, offset)
	reader.Seek(int64(start), io.SeekStart)
	reader.Read(tmp)
	fmt.Println(utf8.ValidString(string(tmp)))
	cc := lz77.Unpack(tmp, 100000000)
	//ioutil.WriteFile("/home/gavin/test.jpeg",tmp,0x777)
	//ioutil.WriteFile("/home/gavin/test1.jpg",cc,0x777)
	fmt.Println(string(cc))
	fmt.Println("record 0 : ", start, offset)
}
