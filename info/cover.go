package info

import (
	"path"
)

type Cover struct {
	FullPathFilename string
	Filename         string
	Ext              string
	TmpDir           string
	Enable           bool
}

func NewCover(fullPathFilename string) *Cover {
	cover := &Cover{
		FullPathFilename: fullPathFilename,
		Filename:         path.Base(fullPathFilename),
		Ext:              path.Ext(fullPathFilename),
		Enable:           true,
	}
	return cover
}
