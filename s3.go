package main

import (
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
	log.Println(CheckBucketExists(sess, "hacker-news-data-cdp"))
	FileUploader(sess, "example1.bleve/index_meta.json")

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
	return false, nil
}

// FileUploader uploads
func FileUploader(sess *session.Session, fp string) error {

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
