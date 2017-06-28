package lz77

import (
	"bytes"
	"fmt"
	"io"
)

const (
	DATA_LENGTH            = 4096
	MAX_SCROLL_WINDOW_SIZE = 2047
	MIN_SCROLL_WINDOW_SIZE = 3

	MAX_FORWARD_WINDOW_SIZE = 10
	MIN_FORWARD_WINDOW_SIZE = 3

	MAX_PLAIN_DATA_SIZE = 4096
	MIN_PLAIN_DATA_SIZE = 0

	MAX_COUNT = 0x08

	LZ_LITERAL       = 0
	LZ_LEN_DIS_PAIR  = 1
	LZ_BYTE_PAIR     = 2
	LZ_LITERAL_COUNT = 3
)

func Unpack(compressed_data []byte, uncompressed_data_max_length int) []byte {
	var uncompressed_data = make([]byte, 0)
	var p = 0

	for p < len(compressed_data) {
		if len(uncompressed_data) >= uncompressed_data_max_length {
			return uncompressed_data
		}
		switch get_byte_type(compressed_data[p]) {
		case LZ_LITERAL:
			uncompressed_data = append(uncompressed_data, compressed_data[p])
			break

		case LZ_LEN_DIS_PAIR:
			var distance = get_distance(compressed_data[p : p+2])
			var length = get_length(compressed_data[p+1])

			var start = len(uncompressed_data) - distance

			//fmt.Println(len(uncompressed_data),start,length,start+length)
			tmpCopy := uncompressed_data[start : start+length]
			uncompressed_data = append(uncompressed_data, tmpCopy...)

			p = p + 1
			break
		case LZ_BYTE_PAIR:
			uncompressed_data = append(uncompressed_data, ' ', compressed_data[p]^128)
			break
		case LZ_LITERAL_COUNT:
			uncompressed_data = append(uncompressed_data, compressed_data[p+1:p+1+int(compressed_data[p])]...)
			p = p + int(compressed_data[p])
			break
		default:
			break
		}
		p++
	}

	return uncompressed_data
}

func get_byte_type(b byte) int {
	if b <= 0xbf && b >= 0x80 {
		return LZ_LEN_DIS_PAIR
	} else if b >= 0xc0 {
		return LZ_BYTE_PAIR
	} else if b >= 0x01 && b <= 0x08 {
		return LZ_LITERAL_COUNT
	} else {
		return LZ_LITERAL
	}
}

func get_distance(b []byte) int {
	var distance uint = 0
	distance = uint(b[0]) & 0x3f
	distance = (distance << 5) + ((uint(b[1]) & 0xf8) >> 3)
	return int(distance)
}

func get_length(b byte) int {
	var length uint = 0
	length = (uint(b) & 0x07) + 3
	return int(length)
}

func DecompressByte(b []byte) []byte {
	reader := bytes.NewReader(b)
	var result = make([]byte, 0)
	for {
		if tmpb, err := reader.ReadByte(); err != io.EOF {
			if tmpb >= 192 {
				result = append(result, 32)
				fmt.Println(^tmpb | 0x80)
				result = append(result, ^tmpb|0x80)
			} else if tmpb >= 128 && tmpb <= 191 {
				//var tmp6 uint16
				//tmp6 = uint16(tmpb >> 2)
				//tmpnext,_ := reader.ReadByte()
				//var tmp14 uint16
				//tmp14 = tmp6 << 8
				//tmp14 = tmp14 | uint16(tmpnext)
				//tmp11 := tmp14>>3
				//tmp3 := tmp14 & 7
				//
				//tmpResult := result[len(result)-int(tmp11):tmp3+3]
				//result = append(result,tmpResult...)
				result = append(result, tmpb)

			} else {
				result = append(result, tmpb)
			}
		} else {
			return result
		}

	}
}
