package utils

import (
	"bytes"
	"encoding/binary"
	"io"
	"strconv"
)

type MOBIHeader struct {
	Compression                      uint16
	Unused0                          uint16
	TextLength                       uint32 //Uncompressed length of the entire text of the book
	RecordCount                      uint16 //Number of PDB records used for the text of the book.
	RecordSize                       uint16 //Maximum size of each record containing text, always 4096
	EncryptionType                   uint16 //0 == no encryption, 1 = Old Mobipocket Encryption, 2 = Mobipocket Encryption
	Unused1                          uint16 //Usually zero
	Identifier                       [4]byte
	HeaderLength                     uint32 // from offset 0x10
	MobiType                         uint32
	TextEncoding                     uint32 //1252 = CP1252 (WinLatin1); 65001 = UTF-8
	UniqueID                         uint32 //Some kind of unique ID number (random?)
	FileVersion                      uint32 //Version of the Mobipocket format used in this file.
	OrthographicIndex                uint32 // 	Section number of orthographic meta index. 0xFFFFFFFF if index is not available.
	InflectionIndex                  uint32 //Section number of inflection meta index. 0xFFFFFFFF if index is not available.
	IndexNames                       uint32 //0xFFFFFFFF if index is not available.
	IndexKeys                        uint32 //0xFFFFFFFF if index is not available.
	ExtraIndex0                      uint32 //Section number of extra 0 meta index. 0xFFFFFFFF if index is not available.
	ExtraIndex1                      uint32
	ExtraIndex2                      uint32
	ExtraIndex3                      uint32
	ExtraIndex4                      uint32
	ExtraIndex5                      uint32
	FirstNonBookIndex                uint32 //First record number (starting with 0) that's not the book's text
	FullNameOffset                   uint32 //Offset in record 0 (not from start of file) of the full name of the book
	FullNameLength                   uint32 //Length in bytes of the full name of the book
	Locale                           uint32 //Book locale code. Low byte is main language 09= English, next byte is dialect, 08 = British, 04 = US. Thus US English is 1033, UK English is 2057.
	InputLanguage                    uint32 //Input language for a dictionary
	OutputLanguage                   uint32 //Output language for a dictionary
	MinVersion                       uint32 //Minimum mobipocket version support needed to read this file.
	FirstImageIndex                  uint32 //First record number (starting with 0) that contains an image. Image records should be sequential.
	HuffmanRecordOffset              uint32 //The record number of the first huffman compression record.
	HuffmanRecordCount               uint32
	HuffmanTableOffset               uint32
	HuffmanTableLength               uint32
	ExthFlags                        uint32   //bitfield. if bit 6 (0x40) is set, then there's an EXTH record
	Unknown32                        [32]byte //32 unknown bytes, if MOBI is long enough
	Unknown                          uint32   //Use 0xFFFFFFFF
	DRMOffset                        uint32   //Offset to DRM key info in DRMed files. 0xFFFFFFFF if no DRM
	DRMCount                         uint32   //Number of entries in DRM info. 0xFFFFFFFF if no DRM
	DRMSize                          uint32   //Number of bytes in DRM info.
	DRMFlags                         uint32   //Some flags concerning the DRM info.
	Unknown8                         [8]byte  //Bytes to the end of the MOBI header, including the following if the header length >= 228 (244 from start of record). Use 0x0000000000000000.
	FirstContentRecordNumber         uint16   //Number of first text record. Normally 1.
	LastContentRecordNumber          uint16   //Number of last image record or number of last text record if it contains no images. Includes Image, DATP, HUFF, DRM.
	Unknown1                         uint32   //Use 0x00000001.
	FCISRecordNumber                 uint32
	Unknown2                         uint32 //Unknown (FCIS record count?) 	Use 0x00000001.
	FLISRecordNumber                 uint32
	Unknown3                         uint32 //Unknown (FLIS record count?) 	Use 0x00000001.
	Unknown4                         uint64 //Use 0x0000000000000000.
	Unknown5                         uint32 //Use 0xFFFFFFFF.
	FirstCompilationDataSectionCount uint32 //Use 0x00000000.
	NumberOfCompilationDataSections  uint32 //Use 0xFFFFFFFF.
	Unknown6                         uint32 //Use 0xFFFFFFFF.
	ExtraRecordDataFlags             uint32 //A set of binary flags, some of which indicate extra data at the end of each text block. This only seems to be valid for Mobipocket format version 5 and 6 (and higher?), when the header length is 228 (0xE4) or 232 (0xE8).
	INDXRecordOffset                 uint32 //(If not 0xFFFFFFFF)The record number of the first INDX record created from an ncx file.
	UnknowUnion24                    [24]byte

	EXTHHeader *EXTHHeader
	FullName   string
}

func NewMOBIHeader() *MOBIHeader {
	header := &MOBIHeader{}
	return header
}

func (header *MOBIHeader) Parse(reader *bytes.Reader, start int64, mobiHeaderSize int) error {
	reader.Seek(start, io.SeekStart)
	binary.Read(reader, binary.BigEndian, &header.Compression)
	binary.Read(reader, binary.BigEndian, &header.Unused0)
	binary.Read(reader, binary.BigEndian, &header.TextLength)
	binary.Read(reader, binary.BigEndian, &header.RecordCount)
	binary.Read(reader, binary.BigEndian, &header.RecordSize)
	binary.Read(reader, binary.BigEndian, &header.EncryptionType)
	binary.Read(reader, binary.BigEndian, &header.Unused1)
	binary.Read(reader, binary.BigEndian, &header.Identifier)
	binary.Read(reader, binary.BigEndian, &header.HeaderLength)
	binary.Read(reader, binary.BigEndian, &header.MobiType)
	binary.Read(reader, binary.BigEndian, &header.TextEncoding)
	binary.Read(reader, binary.BigEndian, &header.UniqueID)
	binary.Read(reader, binary.BigEndian, &header.FileVersion)
	binary.Read(reader, binary.BigEndian, &header.OrthographicIndex)
	binary.Read(reader, binary.BigEndian, &header.InflectionIndex)
	binary.Read(reader, binary.BigEndian, &header.IndexNames)
	binary.Read(reader, binary.BigEndian, &header.IndexKeys)
	binary.Read(reader, binary.BigEndian, &header.ExtraIndex0)
	binary.Read(reader, binary.BigEndian, &header.ExtraIndex1)
	binary.Read(reader, binary.BigEndian, &header.ExtraIndex2)
	binary.Read(reader, binary.BigEndian, &header.ExtraIndex3)
	binary.Read(reader, binary.BigEndian, &header.ExtraIndex4)
	binary.Read(reader, binary.BigEndian, &header.ExtraIndex5)
	binary.Read(reader, binary.BigEndian, &header.FirstNonBookIndex)
	binary.Read(reader, binary.BigEndian, &header.FullNameOffset)
	binary.Read(reader, binary.BigEndian, &header.FullNameLength)
	binary.Read(reader, binary.BigEndian, &header.Locale)
	binary.Read(reader, binary.BigEndian, &header.InputLanguage)
	binary.Read(reader, binary.BigEndian, &header.OutputLanguage)
	binary.Read(reader, binary.BigEndian, &header.MinVersion)
	binary.Read(reader, binary.BigEndian, &header.FirstImageIndex)
	binary.Read(reader, binary.BigEndian, &header.HuffmanRecordOffset)
	binary.Read(reader, binary.BigEndian, &header.HuffmanRecordCount)
	binary.Read(reader, binary.BigEndian, &header.HuffmanTableOffset)
	binary.Read(reader, binary.BigEndian, &header.HuffmanTableLength)
	binary.Read(reader, binary.BigEndian, &header.ExthFlags)
	binary.Read(reader, binary.BigEndian, &header.Unknown32)
	binary.Read(reader, binary.BigEndian, &header.Unknown)
	binary.Read(reader, binary.BigEndian, &header.DRMOffset)
	binary.Read(reader, binary.BigEndian, &header.DRMCount)
	binary.Read(reader, binary.BigEndian, &header.DRMSize)
	binary.Read(reader, binary.BigEndian, &header.DRMFlags)
	binary.Read(reader, binary.BigEndian, &header.Unknown8)
	binary.Read(reader, binary.BigEndian, &header.FirstContentRecordNumber)
	binary.Read(reader, binary.BigEndian, &header.LastContentRecordNumber)
	binary.Read(reader, binary.BigEndian, &header.Unknown1)
	binary.Read(reader, binary.BigEndian, &header.FCISRecordNumber)
	binary.Read(reader, binary.BigEndian, &header.Unknown2)
	binary.Read(reader, binary.BigEndian, &header.FLISRecordNumber)
	binary.Read(reader, binary.BigEndian, &header.Unknown3)
	binary.Read(reader, binary.BigEndian, &header.Unknown4)
	binary.Read(reader, binary.BigEndian, &header.Unknown5)
	binary.Read(reader, binary.BigEndian, &header.FirstCompilationDataSectionCount)
	binary.Read(reader, binary.BigEndian, &header.NumberOfCompilationDataSections)
	binary.Read(reader, binary.BigEndian, &header.Unknown6)
	binary.Read(reader, binary.BigEndian, &header.ExtraRecordDataFlags)
	binary.Read(reader, binary.BigEndian, &header.INDXRecordOffset)
	binary.Read(reader, binary.BigEndian, &header.UnknowUnion24)

	reader.Seek(start+int64(header.FullNameOffset), io.SeekStart)
	tmp := make([]byte, header.FullNameLength)
	reader.Read(tmp)
	header.FullName = string(tmp)

	var exthExists = (header.ExthFlags & 0x40) != 0
	if exthExists {
		header.EXTHHeader = NewEXTHHeader()
		if err := header.EXTHHeader.Parse(reader, start+248); err != nil {
			return err
		}
	}
	return nil
}

func (header *MOBIHeader) GetCharacterEncoding() string {
	switch header.TextEncoding {
	case 1252:
		return "Cp1252"
		break
	case 65001:
		return "UTF-8"
		break
	default:
		break
	}
	return ""
}

func (header *MOBIHeader) exthHeaderSize() int {
	if header.EXTHHeader != nil {
		return header.EXTHHeader.Size()
	}
	return 0
}

func (header *MOBIHeader) getMobiType() string {
	if header.MobiType == 2 {
		return "Mobipocket Book"
	} else if header.MobiType == 3 {
		return "PalmDoc Book"
	} else if header.MobiType == 4 {
		return "Audio"
	} else if header.MobiType == 232 {
		return "mobipocket? generated by kindlegen1.2"
	} else if header.MobiType == 248 {
		return " KF8: generated by kindlegen2"
	} else if header.MobiType == 257 {
		return "News"
	} else if header.MobiType == 258 {
		return "News Feed"
	} else if header.MobiType == 259 {
		return "News Magazine"
	} else if header.MobiType == 513 {
		return "PICS"
	} else if header.MobiType == 514 {
		return "WORD"
	} else if header.MobiType == 515 {
		return "XLS"
	} else if header.MobiType == 516 {
		return "PPT"
	} else if header.MobiType == 517 {
		return "TEXT"
	} else if header.MobiType == 518 {
		return "HTML"
	} else {
		return "Unknown (" + strconv.Itoa(int(header.MobiType)) + ")"
	}
}
func (header *MOBIHeader) getCompression() string {
	switch header.Compression {
	case 1:
		return "None"
		break
	case 2:
		return "PalmDOC"
		break
	case 17480:
		return "HUFF/CDIC"
		break
	default:
		return "Unknown (" + strconv.Itoa(int(header.Compression)) + ")"
	}
	return ""
}

func (header *MOBIHeader) getEncryptionType() string {
	switch header.EncryptionType {
	case 0:
		return "None"
		break
	case 1:
		return "Old Mobipocket"
		break
	case 2:
		return "Mobipocket"
		break
	default:
		return "Unknown (" + strconv.Itoa(int(header.EncryptionType)) + ")"
		break
	}
	return ""
}
