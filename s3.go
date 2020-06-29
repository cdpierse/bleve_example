package main

import (
	"log"
	_ "log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
	log.Println(CheckBucketExists(sess,"hacker-news-data-cdp"))

}

func NewSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(S3REGION)},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil

}

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
