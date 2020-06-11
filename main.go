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
	res, err := ReadDataset(*dataset)
	if err != nil {
		log.Print(err)
	}
	log.Println(res[0].Authors)

}
