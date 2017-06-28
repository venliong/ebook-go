package utils

import (
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
)

type MobiMeta struct {
	PDBHeader  *PDBHeader
	MOBIHeader *MOBIHeader

	cover []byte
}

func NewMobiMeta() *MobiMeta {
	meta := &MobiMeta{
		PDBHeader:  NewPDBHeader(),
		MOBIHeader: NewMOBIHeader(),
		cover:      make([]byte, 0),
	}
	return meta
}

func (meta *MobiMeta) Parse(b []byte) error {
	reader := bytes.NewReader(b)
	var err error
	if err = meta.PDBHeader.Parse(reader); err == nil {
		start, offset := meta.PDBHeader.GetRecord(0)
		if err = meta.MOBIHeader.Parse(reader, int64(start), int(offset)); err == nil {
			if meta.MOBIHeader.FirstImageIndex > 0 && int(meta.MOBIHeader.FirstImageIndex) < len(meta.PDBHeader.RecordInfos) {
				//has cover
				start, offset := meta.PDBHeader.GetRecord(1)
				tmp := make([]byte, offset)
				reader.Seek(int64(start), io.SeekStart)
				reader.Read(tmp)
				meta.cover = tmp
			}
			if meta.MOBIHeader.getCompression() == "PalmDOC" {
				return nil
			} else {
				return errors.New("not a mobi")
			}
		}
	}
	return err
}

func (meta *MobiMeta) GetCharacterEncoding() string {
	return meta.MOBIHeader.GetCharacterEncoding()
}

func (meta *MobiMeta) GetFullName() string {
	return meta.MOBIHeader.FullName
}

func (meta *MobiMeta) GetCover() []byte {
	return meta.cover
}

func (meta *MobiMeta) GetMetaInfo() string {
	var info = ""
	info += "PDB Header\r\n"
	info += "----------\r\n"
	info += "Name: "
	info += meta.PDBHeader.GetName()
	info += "\r\n"
	if len(meta.getPDBHeaderAttributes()) > 0 {
		info += "Attributes: "
		info += strings.Join(meta.getPDBHeaderAttributes(), ",")
		info += "\r\n"
	}
	info += "Version: "
	info += strconv.Itoa(int(meta.PDBHeader.GetVersion()))
	info += "\r\n"
	info += "Creation Date: "
	info += strconv.Itoa(int(meta.PDBHeader.GetCreationDate()))
	info += "\r\n"
	info += "Modification Date: "
	info += strconv.Itoa(int(meta.PDBHeader.GetModificationDate()))
	info += "\r\n"
	info += "Last Backup Date: "
	info += strconv.Itoa(int(meta.PDBHeader.GetLastBackupDate()))
	info += "\r\n"
	info += "Modification Number: "
	info += strconv.Itoa(int(meta.PDBHeader.GetModificationNumber()))
	info += "\r\n"
	info += "App Info ID: "
	info += strconv.Itoa(int(meta.PDBHeader.GetAppInfoID()))
	info += "\r\n"
	info += "Sort Info ID: "
	info += strconv.Itoa(int(meta.PDBHeader.GetSortInfoID()))
	info += "\r\n"
	info += "Type: "
	info += strconv.Itoa(int(meta.PDBHeader.GetType()))
	info += "\r\n"
	info += "Creator: "
	info += strconv.Itoa(int(meta.PDBHeader.GetCreator()))
	info += "\r\n"
	info += "Unique ID Seed: "
	info += strconv.Itoa(int(meta.PDBHeader.GetUniqueIDSeed()))
	info += "\r\n\r\n"

	info += "PalmDOC Header\r\n"
	info += "--------------\r\n"
	info += "Compression: "
	info += meta.MOBIHeader.getCompression()
	info += "\r\n"
	info += "Text Length: "
	info += strconv.Itoa(int(meta.MOBIHeader.TextLength))
	info += "\r\n"
	info += "Record Count: "
	info += strconv.Itoa(int(meta.MOBIHeader.RecordCount))
	info += "\r\n"
	info += "Record Size: "
	info += strconv.Itoa(int(meta.MOBIHeader.RecordSize))
	info += "\r\n"
	info += "Encryption Type: "
	info += meta.MOBIHeader.getEncryptionType()
	info += "\r\n\r\n"

	info += "MOBI Header\r\n"
	info += "-----------\r\n"
	info += "Header Length: "
	info += strconv.Itoa(int(meta.MOBIHeader.HeaderLength))
	info += "\r\n"
	info += "Mobi Type: "
	info += meta.MOBIHeader.getMobiType()
	info += "\r\n"
	info += "Unique ID: "
	info += strconv.Itoa(int(meta.MOBIHeader.UniqueID))
	info += "\r\n"
	info += "File Version: "
	info += strconv.Itoa(int(meta.MOBIHeader.FileVersion))
	info += "\r\n"
	info += "Orthographic Index: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.OrthographicIndex), 16)
	info += "\r\n"
	info += "Inflection Index: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.InflectionIndex), 16)
	info += "\r\n"
	info += "Index Names: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.IndexNames), 16)
	info += "\r\n"
	info += "Index Keys: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.IndexKeys), 16)
	info += "\r\n"
	info += "Extra Index 0: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.ExtraIndex0), 16)
	info += "\r\n"
	info += "Extra Index 1: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.ExtraIndex1), 16)
	info += "\r\n"
	info += "Extra Index 2: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.ExtraIndex2), 16)
	info += "\r\n"
	info += "Extra Index 3: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.ExtraIndex3), 16)
	info += "\r\n"
	info += "Extra Index 4: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.ExtraIndex4), 16)
	info += "\r\n"
	info += "Extra Index 5: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.ExtraIndex5), 16)
	info += "\r\n"
	info += "First Non-Book Index: "
	info += strconv.Itoa(int(meta.MOBIHeader.FirstNonBookIndex))
	info += "\r\n"
	info += "First Image Index: "
	info += strconv.Itoa(int(meta.MOBIHeader.FirstImageIndex))
	info += "\r\n"
	info += "Full Name Offset: "
	info += strconv.Itoa(int(meta.MOBIHeader.FullNameOffset))
	info += "\r\n"
	info += "Full Name Length: "
	info += strconv.Itoa(int(meta.MOBIHeader.FullNameLength))
	info += "\r\n"
	info += "Min Version: "
	info += strconv.Itoa(int(meta.MOBIHeader.MinVersion))
	info += "\r\n"
	info += "Huffman Record Offset: "
	info += strconv.Itoa(int(meta.MOBIHeader.HuffmanRecordOffset))
	info += "\r\n"
	info += "Huffman Record Count: "
	info += strconv.Itoa(int(meta.MOBIHeader.HuffmanRecordCount))
	info += "\r\n"
	info += "Huffman Table Offset: "
	info += strconv.Itoa(int(meta.MOBIHeader.HuffmanTableOffset))
	info += "\r\n"
	info += "Huffman Table Length: "
	info += strconv.Itoa(int(meta.MOBIHeader.HuffmanTableLength))
	info += "\r\n"
	info += "INDX Record Offset: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.INDXRecordOffset), 16)
	info += "\r\n"
	info += "FirstContentRecordNumber: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.FirstContentRecordNumber), 10)
	info += "\r\n"
	info += "LastContentRecordNumber: "
	info += strconv.FormatInt(int64(meta.MOBIHeader.LastContentRecordNumber), 10)
	info += "\r\n"

	return info
}

func (meta *MobiMeta) getPDBHeaderAttributes() []string {
	var attrs = make([]string, 0)
	var attr = meta.PDBHeader.GetAttributes()
	if (attr & 0x02) != 0 {
		attrs = append(attrs, "Read-Only")
	}
	if (attr & 0x04) != 0 {
		attrs = append(attrs, "Dirty AppInfoArea")
	}
	if (attr & 0x08) != 0 {
		attrs = append(attrs, "Backup This Database")
	}
	if (attr & 0x10) != 0 {
		attrs = append(attrs, "OK To Install Newer Over Existing Copy")
	}
	if (attr & 0x20) != 0 {
		attrs = append(attrs, "Force The PalmPilot To Reset After This Database Is Installed")
	}
	if (attr & 0x40) != 0 {
		attrs = append(attrs, "Don't Allow Copy Of File To Be Beamed To Other Pilot")
	}
	return attrs
}
