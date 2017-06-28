package utils

import (
	"archive/zip"
	"github.com/google/uuid"
	"io"
	"os"
	"path"
	"path/filepath"
)

type Zip struct {
	FullPathFilename string
	Filename         string
	Ext              string
	TmpDir           string
}

func NewZip(fullPathFilename string) *Zip {
	zip := &Zip{
		FullPathFilename: fullPathFilename,
		Filename:         path.Base(fullPathFilename),
		Ext:              path.Ext(fullPathFilename),
		TmpDir:           path.Join(os.TempDir(), uuid.New().String()),
	}
	return zip
}

func NewZipWithTmpDir(fullPathFilename, tmpDir string) *Zip {
	zip := &Zip{
		FullPathFilename: fullPathFilename,
		Filename:         path.Base(fullPathFilename),
		Ext:              path.Ext(fullPathFilename),
		TmpDir:           tmpDir,
	}
	return zip
}

func (z Zip) Decompress() error {
	if readCloser, err := zip.OpenReader(z.FullPathFilename); err == nil {
		defer readCloser.Close()
		for _, f := range readCloser.File {
			if rc, err := f.Open(); err == nil {
				defer rc.Close()
				os.MkdirAll(filepath.Join(z.TmpDir, path.Dir(f.Name)), 0777)
				var file, _ = os.Create(filepath.Join(z.TmpDir, f.Name))
				if _, err = io.CopyN(file, rc, int64(f.UncompressedSize64)); err != nil {
					return err
				}
			} else {
				return err
			}
		}
		return nil
	} else {
		return err
	}
}

func (z Zip) Close() {
	os.RemoveAll(z.TmpDir)
}
