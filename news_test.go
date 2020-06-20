package main

import (
	"log"
	"os"
	"testing"
)

func setupArticles() Articles {
	articles, err := ReadDataset("News_Category_Dataset_v2.json")
	if err != nil {
		log.Fatal("Could not read news dataset")
	}
	return articles[0:100]

}

func teardown() {
	err := os.RemoveAll("testIndex.bleve/")
	if err != nil {
		log.Fatal("Could not remove test index directory")
	}
}

func TestReadDataset(t *testing.T) {
	name := "News_Category_Dataset_v2.json"
	_, err := ReadDataset(name)
	if err != nil {
		t.Errorf("Could not read from %s correctly", name)
	}
}

func TestNewArticleIndex(t *testing.T) {
	articles := setupArticles()
	err := NewArticleIndex(articles, "testIndex.bleve")
	if err != nil {
		t.Fatal("Could not create test index.")
	}
	log.Println("Got here")
	
	// TODO: #8 @cdpierse fix failing conditions for dir check
	ok := dirExists("testIndex.bleve")
	if !ok {
		t.Errorf("Did not create article index")
	}
	// _, err = GetIndex("testIndex.bleve")
	// if err != nil {
	// 	t.Errorf("Could not load the test index created")
	// }
	teardown()

}
