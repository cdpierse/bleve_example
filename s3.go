package main

import (
	"errors"
	"log"
	_ "log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/blevesearch/bleve"
)

// Change these depending on your bucket, and region
const (
	S3REGION = "eu-west-1"
	S3BUCKET = "bleve-search-example"
)

// UploadIndexDir uploads the given bleve index an S3 bucket
func UploadIndexDir(index *bleve.Index, name string) error {
	return nil
}

// DownloadIndexDir downloads the latest saved version of an index
// from S3
func DownloadIndexDir(name string) error {
	return nil
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(S3REGION)},
	)
	if err != nil {
		panic(err)
	}

	if ok, err := CheckBucketExists(sess, S3BUCKET); !ok {
		panic(err)
	}

	// UploadIndex("example1.bleve")
	resp, err := CheckBucketObjectExists(sess, "index_mta.json", S3BUCKET)
	if err != nil {
		log.Panic(err)
	}
	log.Println(resp)

}

// NewSession creates a new seission service client.
func NewSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(S3REGION)},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil

}

// CheckBucketExists checks if a bucket exists for given aws account
//	Inputs:
//     sess: is the current session, which provides configuration for the SDK's service clients
//     name: is the name of the bucket
//	Output:
//     true if the bucket exists, false if not
//     error will not be nil if something goes wrong
func CheckBucketExists(sess *session.Session, name string) (bool, error) {
	client := s3.New(sess)
	result, err := client.ListBuckets(nil)
	if err != nil {
		return false, err
	}
	for _, bucket := range result.Buckets {
		if *bucket.Name == name {
			return true, nil
		}
	}
	return false, errors.New("s3 bucket '" + name + " ' does not exist.")
}

func CheckBucketObjectExists(sess *session.Session, filename string, bucket string) (bool, error) {
	client := s3.New(sess)
	response, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return false, err
	}
	for _, obj := range response.Contents {
		if filename == *obj.Key {
			return true, nil
		}
	}

	return false, nil
}

// FileUpload uploads a single file found at fp to an s3 bucket
func FileUpload(sess *session.Session, fp string) error {

	file, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	_, filename := filepath.Split(fp)

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3BUCKET),
		Key:    aws.String(filename),
		Body:   file,
	})

	return nil
}

// UploadIndex finds all files in a given directory and individually adds each
// to the bleve index bucket.
func UploadIndex(root string) error {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	sess, err := NewSession()
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		log.Printf("Uploading: %v", file)
		_ = FileUpload(sess, file)
	}
	return nil

}
