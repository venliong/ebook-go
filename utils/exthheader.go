package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	//"fmt"
)

type EXTHHeader struct {
	Identifier   [4]byte
	HeaderLength uint32
	RecordCount  uint32
	EXTHRecords  []*EXTHRecord
}

func NewEXTHHeader() *EXTHHeader {
	header := &EXTHHeader{
		EXTHRecords: make([]*EXTHRecord, 0),
	}
	return header
}

func (header *EXTHHeader) Parse(reader *bytes.Reader, start int64) error {
	reader.Seek(start, io.SeekStart)
	binary.Read(reader, binary.BigEndian, &header.Identifier)
	//var tmp int64 = 0
	//for  tmp < int64(reader.Len()) {
	//	reader.Seek(tmp,io.SeekStart)
	//	binary.Read(reader, binary.BigEndian, &header.Identifier)
	//	if header.Identifier[0] != 69 || header.Identifier[1] != 88 || header.Identifier[2] != 84 || header.Identifier[3] != 72 {
	//		tmp++
	//	}else {
	//		fmt.Println(tmp)
	//		return errors.New("Expected to find EXTH header identifier EXTH but got something else instead")
	//	}
	//}
	if header.Identifier[0] != 69 || header.Identifier[1] != 88 || header.Identifier[2] != 84 || header.Identifier[3] != 72 {
		return errors.New("Expected to find EXTH header identifier EXTH but got something else instead")
	}
	binary.Read(reader, binary.BigEndian, &header.HeaderLength)
	binary.Read(reader, binary.BigEndian, &header.RecordCount)

	header.EXTHRecords = make([]*EXTHRecord, 0)
	for len(header.EXTHRecords) < int(header.RecordCount) {
		record := NewEXTHRecord()
		record.Parse(reader)
		header.EXTHRecords = append(header.EXTHRecords, record)
	}

	var padding = header.paddingSize(header.dataSize())
	var paddingb = make([]byte, padding)
	reader.Read(paddingb)
	return nil
}

func (header *EXTHHeader) Size() int {
	var dataSize = header.dataSize()
	return 12 + dataSize + header.paddingSize(dataSize)
}

func (header *EXTHHeader) dataSize() int {
	var dataSize = 0
	for _, value := range header.EXTHRecords {
		dataSize += value.Size()
	}
	return dataSize
}
func (header *EXTHHeader) paddingSize(dataSize int) int {
	var paddingSize = dataSize % 4
	if paddingSize != 0 {
		paddingSize = 4 - paddingSize
	}
	return paddingSize
}
