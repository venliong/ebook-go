package utils

import (
	"bytes"
	"encoding/binary"
)

type RecordInfo struct {
	RecordDataOffset [4]byte
	RecordAttributes [1]byte
	UniqueID         [3]byte
}

func NewRecordInfo() *RecordInfo {
	info := &RecordInfo{}
	return info
}

func (record *RecordInfo) GetRecordDataOffset() int32 {
	bts := make([]byte, 0)
	if len(record.RecordDataOffset) < 4 {
		var padding = 4 - len(record.RecordDataOffset)
		var data = 0
		for len(bts) < padding {
			bts = append(bts, 0)
			data += data
		}
		data = 0
		for len(bts) < 4 {
			bts = append(bts, record.RecordDataOffset[data])
			data += 1
		}
	} else {
		var data = 0
		for len(bts) < 4 {
			bts = append(bts, record.RecordDataOffset[data])
			data += 1
		}
	}
	var result int32
	binary.Read(bytes.NewReader(bts), binary.BigEndian, &result)
	return result
}

func (record *RecordInfo) GetUniqueID() int32 {
	bts := make([]byte, 0)
	if len(record.UniqueID) < 4 {
		var padding = 4 - len(record.UniqueID)
		var data = 0
		for len(bts) < padding {
			bts = append(bts, 0)
			data += data
		}
		data = 0
		for len(bts) < 4 {
			bts = append(bts, record.UniqueID[data])
			data += 1
		}
	} else {
		var data = 0
		for len(bts) < 4 {
			bts = append(bts, record.UniqueID[data])
			data += 1
		}
	}
	var result int32
	binary.Read(bytes.NewReader(bts), binary.BigEndian, &result)
	return result
}
