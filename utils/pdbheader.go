package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type PDBHeader struct {
	Name               [32]byte
	Attributes         int16
	Version            int16
	CreationDate       int32
	ModificationDate   int32
	LastBackupDate     int32
	ModificationNumber int32
	AppInfoID          int32
	SortInfoID         int32
	Type               int32
	Creator            int32
	UniqueIDSeed       int32
	NextRecordListID   int32
	NumRecords         int16

	RecordInfos []*RecordInfo

	GapToData int16
}

func NewPDBHeader() *PDBHeader {
	header := &PDBHeader{
		RecordInfos: make([]*RecordInfo, 0),
	}
	return header
}

func (header *PDBHeader) Parse(reader *bytes.Reader) error {
	if reader.Len() < 78 {
		return errors.New("not a pdb")
	}
	reader.Seek(76, io.SeekStart)
	binary.Read(reader, binary.BigEndian, &header.NumRecords)

	for len(header.RecordInfos) < int(header.NumRecords) {
		record := NewRecordInfo()
		header.RecordInfos = append(header.RecordInfos, record)
	}
	reader.Seek(0, io.SeekStart)
	binary.Read(reader, binary.BigEndian, &header.Name)
	binary.Read(reader, binary.BigEndian, &header.Attributes)
	binary.Read(reader, binary.BigEndian, &header.Version)
	binary.Read(reader, binary.BigEndian, &header.CreationDate)
	binary.Read(reader, binary.BigEndian, &header.ModificationDate)
	binary.Read(reader, binary.BigEndian, &header.LastBackupDate)
	binary.Read(reader, binary.BigEndian, &header.ModificationNumber)
	binary.Read(reader, binary.BigEndian, &header.AppInfoID)
	binary.Read(reader, binary.BigEndian, &header.SortInfoID)
	binary.Read(reader, binary.BigEndian, &header.Type)
	binary.Read(reader, binary.BigEndian, &header.Creator)
	binary.Read(reader, binary.BigEndian, &header.UniqueIDSeed)
	binary.Read(reader, binary.BigEndian, &header.NextRecordListID)
	binary.Read(reader, binary.BigEndian, &header.NumRecords)
	for _, value := range header.RecordInfos {
		binary.Read(reader, binary.BigEndian, value)
	}
	binary.Read(reader, binary.BigEndian, &header.GapToData)

	return nil
}

func (header *PDBHeader) GetMobiHeaderSize() int32 {
	if len(header.RecordInfos) > 1 {
		return header.RecordInfos[1].GetRecordDataOffset() - header.RecordInfos[0].GetRecordDataOffset()
	} else {
		return 0
	}
}

func (header *PDBHeader) GetRecord(record int) (start, length int32) {
	if len(header.RecordInfos) > record+1 {
		return header.RecordInfos[record].GetRecordDataOffset(), header.RecordInfos[record+1].GetRecordDataOffset() - header.RecordInfos[record].GetRecordDataOffset()
	} else {
		return 0, 0
	}
}

func (header *PDBHeader) GetNumRecords() int16 {
	return header.NumRecords
}

func (header *PDBHeader) GetName() string {
	var i = 0
	var tmp = make([]byte, 0)
	for i < len(header.Name) {
		tmp = append(tmp, header.Name[i])
		i++
	}
	return string(tmp)
}

func (header *PDBHeader) GetVersion() int16 {
	return header.Version
}

func (header *PDBHeader) GetCreationDate() int32 {
	return header.CreationDate
}

func (header *PDBHeader) GetModificationDate() int32 {
	return header.ModificationDate
}

func (header *PDBHeader) GetLastBackupDate() int32 {
	return header.LastBackupDate
}
func (header *PDBHeader) GetModificationNumber() int32 {
	return header.ModificationNumber
}

//appInfoID
func (header *PDBHeader) GetAppInfoID() int32 {
	return header.AppInfoID
}
func (header *PDBHeader) GetSortInfoID() int32 {
	return header.SortInfoID
}
func (header *PDBHeader) GetType() int32 {
	return header.Type
}
func (header *PDBHeader) GetCreator() int32 {
	return header.Creator
}

func (header *PDBHeader) GetUniqueIDSeed() int32 {
	return header.UniqueIDSeed
}
func (header *PDBHeader) GetAttributes() int16 {
	return header.Attributes
}

func (header *PDBHeader) GetAttributesMean() string {
	//0x0002 Read-Only
	//0x0004 Dirty AppInfoArea
	//0x0008 Backup this database (i.e. no conduit exists)
	//0x0010 (16 decimal) Okay to install newer over existing copy, if present on PalmPilot
	//0x0020 (32 decimal) Force the PalmPilot to reset after this database is installed
	//0x0040 (64 decimal) Don't allow copy of file to be beamed to other Pilot.
	switch header.Attributes {
	case 0x0002:
		return "Read-Only"
	case 0x0004:
		return "Dirty AppInfoArea"
	case 0x0008:
		return "Backup this database"
	case 0x0010:
		return "(16 decimal) Okay to install newer over existing copy, if present on PalmPilot"
	case 0x0020:
		return "0x0020 (32 decimal) Force the PalmPilot to reset after this database is installed"
	case 0x0040:
		return "(64 decimal) Don't allow copy of file to be beamed to other Pilot."
	default:
		return "unknow mean"
	}
}

func (header *PDBHeader) Size() int {
	return 80 + len(header.RecordInfos)*8
}
