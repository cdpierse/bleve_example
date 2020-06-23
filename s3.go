package main

import (
	_"log"

	"github.com/blevesearch/bleve"
)

// UploadIndex uploads the given bleve index an S3 bucket
func UploadIndex(index *bleve.Index, name string) error {
	return nil
}

// DownloadIndex downloads the latest saved version of an index
// from S3
func DownloadIndex(name string) error {
	return nil
}