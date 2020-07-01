package main

import (
	_ "log"
	_ "os"
	"testing"
)

const s3TestIndex = "s3Test.bleve"

// func TestDownload(t *testing.T) {
// 	err := DownloadIndex(s3TestIndex)
// 	if err != nil {
// 		t.Fatal("Error running download")
// 	}

// 	// test if directory for index exists
// 	_, err = os.Stat(s3TestIndex)
// 	if os.IsNotExist(err) {
// 		t.Fatal("Directory does not exist")
// 	}

// 	// test if the standard files created by bleve exist in directory
// 	files := []string{"store", "index_meta.json"}
// 	for _, file := range files {
// 		_, err := os.Stat(s3TestIndex + "/" + file)
// 		if os.IsNotExist(err) {
// 			t.Fatalf("File '%s' does not exist in index", file)
// 		}
// 	}

// 	// test if the index can be loaded correctly
// 	_, err = GetIndex(s3TestIndex)
// 	if err != nil {
// 		t.Fatal("Could not fetch and load index")
// 	}
// }

func TestUpload(t *testing.T) {
	filename := "test.txt"
	sess, err := NewSession()
	if err != nil {
		t.Fatal(err)
	}
	err = FileUpload(sess, filename)
	if err != nil {
		t.Fatal(err)
	}

	if ok, err := ObjectExists(sess, filename, S3BUCKET); !ok {
		t.Fatal(err)
	}
}
