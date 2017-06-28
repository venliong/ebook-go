package info

import "time"

type Book struct {
	Title        string
	Pubdate      string
	Isbn         string
	Path         string
	HasCover     int
	Timestamp    time.Time
	LastModified time.Time
}
