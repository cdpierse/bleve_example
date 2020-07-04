package main

import (
	_ "log"
	_ "os"
	"path/filepath"
	"testing"
)

const s3TestIndex = "s3Test.bleve"

func s3Setup() {

}

func s3Teardown() {

}

func TestUpload(t *testing.T) {
	path := "data/test.txt"
	sess, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	err = UploadFile(sess, path)
	if err != nil {
		t.Fatal(err)
	}

	_, filename := filepath.Split(path)

	if ok, err := ObjectExists(sess, filename, S3BUCKET); !ok {
		t.Fatal(err)
	}
}

func TestDownload(t *testing.T) {
	dirName := "test.bleve"
	err := DownloadIndex()
	if err != nil {
		t.Fatal(err)
	}

	if !dirExists(dirName) {
		t.Fatalf("Directory %s does not exists ", dirName)
	}
	indexFiles := []string{"index_meta.json", "store"}
	for _, file := range indexFiles {
		if !fileExists(file) {
			t.Fatalf("Could not find the '%s' in directory %s", file, dirName)
		}

	}

}
