package main

import (
	"flag"
	"log"
)

func main() {
	log.Println("Hello World")
	dataset := flag.String(
		"dataset",
		"News_Category_Dataset_v2.json",
		"Dataset JSON file to be read and parsed to struct")

	flag.Parse()
	log.Println(*dataset)
	articles, err := ReadDataset(*dataset)
	if err != nil {
		log.Print(err)
	}

	index, err := GetIndex("example1.bleve")
	if err != nil {
		log.Print(err)
	}
	log.Println("Loaded Index")

	matches, err := MatchQueryIndex("California fires rage on",index)
	GetQueryHits(matches,articles)


}
