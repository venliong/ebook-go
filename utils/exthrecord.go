package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
)

var (
	CODETYPESTR  = make(map[uint32]string)
	CODETYPEINT  = make(map[uint32]string)
	CODETYPEADDR = make(map[uint32]string)
)

func init() {
	CODETYPESTR[100] = "author"
	CODETYPESTR[101] = "publisher"
	CODETYPESTR[103] = "description"
	CODETYPESTR[104] = "isbn"
	CODETYPESTR[105] = "subject"
	CODETYPESTR[106] = "publishdate"
	CODETYPESTR[108] = "builder"
	CODETYPESTR[109] = "rights"
	CODETYPESTR[113] = "asin"
	CODETYPESTR[503] = "title"
	CODETYPESTR[504] = "asin"
	CODETYPESTR[118] = "retail price"
	CODETYPESTR[119] = "retail price currency"
	CODETYPESTR[200] = "dictionary short name"
	CODETYPESTR[501] = "cdetype"

	CODETYPEADDR[201] = "coveroffset"
	CODETYPEADDR[202] = "thumboffset"
	CODETYPEINT[404] = "text to speech"
}

type EXTHRecord struct {
	RecordType   uint32
	RecordLength uint32
	RecordData   []byte
}

func NewEXTHRecord() *EXTHRecord {
	record := &EXTHRecord{}
	return record
}

func (record *EXTHRecord) Parse(reader *bytes.Reader) error {
	binary.Read(reader, binary.BigEndian, &record.RecordType)
	binary.Read(reader, binary.BigEndian, &record.RecordLength)

	if record.RecordLength < 8 {
		return errors.New("Invalid EXTH record length")
	}

	record.RecordData = make([]byte, record.RecordLength-8)
	reader.Read(record.RecordData)
	return nil
}

func (record *EXTHRecord) Type() string {
	if value, has := CODETYPESTR[record.RecordType]; has {
		return value
	}

	if value, has := CODETYPEINT[record.RecordType]; has {
		return value
	}

	if value, has := CODETYPEADDR[record.RecordType]; has {
		return value
	}
	return strconv.Itoa(int(record.RecordType))
}

func (record *EXTHRecord) Data() string {
	if _, has := CODETYPESTR[record.RecordType]; has {
		return string(record.RecordData)
	}

	if _, has := CODETYPEINT[record.RecordType]; has {
		var tmp = make([]byte, 1)
		tmp[0] = 0
		if len(record.RecordData) == bytes.Count(record.RecordData, tmp) {
			return "enabled"
		} else {
			return "disabled"
		}
	}

	if _, has := CODETYPEADDR[record.RecordType]; has {
		var tmp int32
		binary.Read(bytes.NewReader(record.RecordData), binary.BigEndian, &tmp)
		return strconv.Itoa(int(tmp))
	}
	return strconv.Itoa(int(record.RecordType))
	//return string(record.RecordData)
}

func (record *EXTHRecord) Size() int {
	return len(record.RecordData) + 8
}
