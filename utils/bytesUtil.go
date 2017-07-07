package utils

import (
	"bytes"
	"encoding/binary"
)

func byte2int16(b []byte) int16 {
	bts := make([]byte, 2)
	if len(b) == 0 {
		return 0
	} else if len(b) == 1 {
		bts[0] = 0
		bts[1] = b[0]
	} else if len(b) >= 2 {
		bts[0] = b[0]
		bts[1] = b[1]
	}
	var result int16
	binary.Read(bytes.NewReader(bts), binary.BigEndian, &result)
	return result
}

func byte2int32(b []byte) int32 {
	bts := make([]byte, 0)
	if len(b) < 4 {
		var padding = 4 - len(b)
		var data = 0
		for len(bts) < padding {
			bts = append(bts, 0)
			data += data
		}
		data = 0
		for len(bts) < 4 {
			bts = append(bts, b[data])
			data += 1
		}
	} else {
		var data = 0
		for len(bts) < 4 {
			bts = append(bts, b[data])
			data += 1
		}
	}
	var result int32
	binary.Read(bytes.NewReader(bts), binary.BigEndian, &result)
	return result
}

func byte2int64(b []byte) int64 {
	bts := make([]byte, 0)
	if len(b) < 8 {
		var padding = 8 - len(b)
		var data = 0
		for len(bts) < padding {
			bts = append(bts, 0)
			data += data
		}
		data = 0
		for len(bts) < 8 {
			bts = append(bts, b[data])
			data += 1
		}
	} else {
		var data = 0
		for len(bts) < 8 {
			bts = append(bts, b[data])
			data += 1
		}
	}
	var result int64
	binary.Read(bytes.NewReader(bts), binary.BigEndian, &result)
	return result
}

func to32Byte(b [32]byte) []byte {
	var tmp = make([]byte, 32)
	for _, value := range b {
		tmp = append(tmp, value)
	}
	return tmp
}
