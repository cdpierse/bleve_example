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
	// articles, err := ReadDataset(*dataset)
	// if err != nil {
	// 	log.Print(err)
	// }

	// index, err := GetIndex("example1.bleve")
	// if err != nil {
	// 	log.Print(err)
	// }
	// log.Println("Loaded Index")

	// longstring := `
	// A combination of dry weather and record heat in recent days has led to several wildfires in California, forcing some evacuations on Wednesday.
	// Officials in Ventura County issued an evacuation order for a remote area of Piru located about 48 miles northwest of downtown Los Angeles for a blaze known as the Lime Fire.
	// The blaze broke out Wednesday afternoon near the Lake Piru campground and has burned about 450 acres so far. 
	// Fire officials said Thursday that two minor injuries have been reported so far, and the fire is 20 percent contained.
	// `
	// matches, err := MatchQueryIndex(longstring, index)
	// GetQueryHits(matches, articles)
	GetTemplate()

}
