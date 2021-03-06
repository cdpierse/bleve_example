package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Hello World")
	dataset := flag.String(
		"dataset",
		"News_Category_Dataset_v2.json",
		"Dataset JSON file to be read and parsed to struct")

	flag.Parse()
	log.Println(*dataset)

	if !fileExists(*dataset) {
		log.Fatalf("No such file %v. Exiting", *dataset)
	}

	articles, err := ReadDataset(*dataset)
	if err != nil {
		log.Print(err)
	}

	if !dirExists("example1.bleve") {
		log.Println("Index does not exist, creating from file:", *dataset)
		NewArticleIndex(articles, "")
	}

	// TODO: #14 @cdpierse Add S3 upload and retrieval functionality flags
	index, err := GetIndex("example1.bleve")
	if err != nil {
		log.Print(err)
	}
	log.Println("Loaded Index")

	tmpl := template.Must(template.ParseFiles("index.html"))
	handler := ServeTemplate(tmpl, articles, index)

	http.ListenAndServe(":3000", handler)

}

func fileExists(name string) bool {
	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()

}

func dirExists(name string) bool {
	info, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}
