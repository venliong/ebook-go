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
	CODETYPETRAN = make(map[uint32]string)
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
	CODETYPESTR[117] = "adult"
	CODETYPESTR[118] = "retail price"
	CODETYPESTR[119] = "retail price currency"
	CODETYPESTR[129] = "KF8 cover URI"
	CODETYPESTR[200] = "dictionary short name"

	CODETYPESTR[300] = "fontsignature"
	CODETYPESTR[501] = "cdetype"
	CODETYPESTR[503] = "title"
	CODETYPESTR[504] = "asin"
	CODETYPESTR[524] = "language"
	CODETYPESTR[525] = "对齐"

	CODETYPEADDR[121] = "KF8 BOUNDARY Offset"
	CODETYPEADDR[125] = "count of resources"
	CODETYPEADDR[131] = "UnknownADDR"
	CODETYPEADDR[201] = "coveroffset"
	CODETYPEADDR[202] = "thumboffset"
	CODETYPEADDR[203] = "hasfakecover"
	CODETYPEADDR[205] = "Creator Major Version"
	CODETYPEADDR[206] = "Creator Minor Version"
	CODETYPEADDR[207] = "Creator Build Number"

	CODETYPEINT[404] = "text to speech"

	CODETYPETRAN[535] = "Creator Build Number"
	CODETYPETRAN[204] = "Creator Software"

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

	if value, has := CODETYPETRAN[record.RecordType]; has {
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

	if _, has := CODETYPETRAN[record.RecordType]; has {
		if record.RecordType == 535 {
			return "a build number of Kindlegen 2.7"
		}
		var tmp int32
		binary.Read(bytes.NewReader(record.RecordData), binary.BigEndian, &tmp)
		switch tmp {
		case 1:
			return "mobigen"
		case 2:
			return "Mobipocket Creator"
		case 200:
			return "kindlegen (Windows)"
		case 201:
			return "kindlegen (Linux)"
		case 202:
			return "kindlegen (Mac)"
		case 33307:
			return "calibre mock Linux kindlegen 1.2"
		case 0:
			return "calibre mock Linux kindlegen 2.0 期刊"
		case 101:
			return "calibre mock Linux kindlegen 2.0 期刊"
		default:
			return strconv.Itoa(int(tmp))
		}

	}

	return strconv.Itoa(int(record.RecordType))
	//return string(record.RecordData)
}

func (record *EXTHRecord) Size() int {
	return len(record.RecordData) + 8
}
