package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"testing"
)

type PDBTestHeader struct {
	Name               [32]byte
	Attributes         [2]byte
	Version            [2]byte
	CreationDate       [4]byte
	ModificationDate   [4]byte
	LastBackupDate     [4]byte
	ModificationNumber [4]byte
	AppInfoID          [4]byte
	SortInfoID         [4]byte
	Type               [4]byte
	Creator            [4]byte
	UniqueIDSeed       [4]byte
	NextRecordListID   [4]byte
	NumRecords         [2]byte
}

func NewPDBTestHeader() *PDBTestHeader {
	header := &PDBTestHeader{
	//Name:               make([]byte, 32),
	//Attributes:         make([]byte, 2),
	//Version:            make([]byte, 2),
	//CreationDate:       make([]byte, 4),
	//ModificationDate:   make([]byte, 4),
	//LastBackupDate:     make([]byte, 4),
	//ModificationNumber: make([]byte, 4),
	//AppInfoID:          make([]byte, 4),
	//SortInfoID:         make([]byte, 4),
	//Type:               make([]byte, 4),
	//Creator:            make([]byte, 4),
	//UniqueIDSeed:       make([]byte, 4),
	//NextRecordListID:   make([]byte, 4),
	//NumRecords:         make([]byte, 2),
	}
	return header
}

func TestPDBHeader_Parse2(t *testing.T) {
	fmt.Println(t.Name())
	header := NewPDBHeader()
	//b, _ := ioutil.ReadFile("../tmp/grdls.mobi")
	//b,_ :=ioutil.ReadFile("/home/gavin/Calibre 书库/Wang Ceng Qi/Shou Jie (3)/sj.mobi")
	b, _ := ioutil.ReadFile("../tmp/st1.azw3")
	reader := bytes.NewReader(b)
	header.Parse(reader)
	fmt.Println(header.Name)
	fmt.Println(header.NumRecords)

	theader := NewPDBTestHeader()
	var tmpb = make([]byte, 78)
	var i = 0
	for i < 78 {
		tmpb[i] = b[i]
		i++
	}
	err := binary.Read(bytes.NewReader(tmpb), binary.BigEndian, theader)
	fmt.Println(err)
	fmt.Println(theader.Name)
	fmt.Println(theader.NumRecords)
}
