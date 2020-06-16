package main

import (
	"flag"
	"log"
	"os"
)

// TODO: #4 Add file exists check and handling
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
		NewArticleIndex(articles)
	}

	index, err := GetIndex("example1.bleve")
	if err != nil {
		log.Print(err)
	}
	log.Println("Loaded Index")

	// longstring := `
	// A combination of dry weather and record heat in recent days has led to several wildfires in California, forcing some evacuations on Wednesday.
	// Officials in Ventura County issued an evacuation order for a remote area of Piru located about 48 miles northwest of downtown Los Angeles for a blaze known as the Lime Fire.
	// The blaze broke out Wednesday afternoon near the Lake Piru campground and has burned about 450 acres so far.
	// Fire officials said Thursday that two minor injuries have been reported so far, and the fire is 20 percent contained.
	// `

	// testString := "conversatian"
	// matches, err := MatchQuery(testString, index)
	matches, err := TermQuery("Igor",index)
	GetQueryHits(matches, articles)
	// GetTemplate()

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
