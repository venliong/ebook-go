package utils

import (
	"bytes"
	"fmt"
	"github.com/jiangmitiao/ebook-go/utils/lz77"
	"io"
	"io/ioutil"
	"testing"
)

func TestMobiMeta_Parse(t *testing.T) {
	fmt.Println(t.Name())

	b, _ := ioutil.ReadFile("../tmp/grdls.mobi")
	//b, _ := ioutil.ReadFile("../tmp/zcdz.mobi")
	//b, _ := ioutil.ReadFile("../tmp/st1.azw3")
	reader := bytes.NewReader(b)
	meta := NewMobiMeta()
	meta.Parse(b)

	fmt.Println(meta.GetCharacterEncoding())
	fmt.Println(meta.GetFullName())
	fmt.Println(meta.GetMetaInfo())

	for _, value := range meta.PDBHeader.RecordInfos {
		fmt.Print(value.GetUniqueID())
	}
	fmt.Println()

	for _, value := range meta.MOBIHeader.EXTHHeader.EXTHRecords {
		fmt.Println(value.Type(), value.Data())
	}

	allb := make([]byte, 0)

	var p = 1
	for p < int(meta.MOBIHeader.FirstNonBookIndex) {
		start, length := meta.PDBHeader.GetRecord(p)
		tmp := make([]byte, length)
		reader.Seek(int64(start), io.SeekStart)
		reader.Read(tmp)
		allb = append(allb, tmp...)
		p++
	}
	cc := lz77.Unpack(allb, 100000000000)
	fmt.Println(string(cc))
	//fmt.Println(string(allb))

	//fmt.Println()
	//start, length := meta.PDBHeader.GetRecord(11)
	//reader.Seek(int64(start), io.SeekStart)
	//tmp := make([]byte, length)
	//reader.Read(tmp)
	////cc=lz77.Unpack(tmp,100000000000)
	//fmt.Println(string(tmp))
}
