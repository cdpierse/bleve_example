package main

import (
	"errors"
	"log"
	_ "log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Change these depending on your bucket, and region
const (
	S3REGION = "eu-west-1"
	S3BUCKET = "bleve-search-example"
)

// NewSession creates a new session service client.
func NewSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(S3REGION)},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil

}

// BucketExists checks if a bucket exists for given aws account
//	Inputs:
//     sess: is the current session, which provides configuration for the SDK's service clients
//     name: is the name of the bucket
//	Output:
//     true if the bucket exists, false if not
//     error will not be nil if something goes wrong
func BucketExists(sess *session.Session, name string) (bool, error) {
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

// ObjectExists asserts whether a file(object) exists in a given bucket.
//	Inputs:
//     sess: is the current session, which provides configuration for the SDK's service clients
//     filename: is the file(object) we are testing if it exists
//	   bucket: is the S3 bucket the filename for in
//	Output:
//     true: if the file exists, false if not
//     error will not be nil if something goes wrong
func ObjectExists(sess *session.Session, filename string, bucket string) (bool, error) {
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
func UploadFile(sess *session.Session, fp string) error {

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
// one at a time to the bleve index bucket.
func UploadIndex(sess *session.Session, root string) error {
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

	if err != nil {
		panic(err)
	}
	for _, file := range files {
		log.Printf("Uploading: %v", file)
		_ = UploadFile(sess, file)
	}
	return nil

}

// todo: #18 @cdpierse add functions to download index from s3

// DownloadFile does stuff downloads a file given by filename from the
// s3 instance given by sess and the S3BUCKET inside that instance.
//	Inputs:
//     sess: is the current session, which provides configuration for the SDK's service clients
//     filename: is the name of the file to be downloaded.
//	Output:
//     error
func DownloadFile(sess *session.Session, filename string) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	downloader := s3manager.NewDownloader(sess)
	requestInput := s3.GetObjectInput{
		Bucket: aws.String(S3BUCKET),
		Key:    aws.String(filename),
	}

	buff := &aws.WriteAtBuffer{}
	_, err = downloader.Download(buff, &requestInput)
	if err != nil {
		return err
	}
	numBytes, err := file.Write(buff.Bytes())
	if err != nil {
		return err
	}
	log.Println(numBytes)

	return nil

}

func DownloadIndex(dirname string) error {
	if !strings.HasSuffix(dirname, ".bleve") {
		dirname = dirname + ".bleve"
	}
	ok := dirExists(dirname)
	if ok {
		log.Printf("Dirname '%s' already exists, exiting function.", dirname)
		return nil
	}
	os.Mkdir(dirname, 0700)
	return nil

}
