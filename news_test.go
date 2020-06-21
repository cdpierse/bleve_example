package main

import (
	_ "github.com/blevesearch/bleve"
	"log"
	"os"
	"testing"
)

const TestIndexName = "testIndex.bleve"

func setupArticles() Articles {
	articles, err := ReadDataset("News_Category_Dataset_v2.json")
	if err != nil {
		log.Fatal("Could not read news dataset")
	}
	return articles[0:100]

}

func teardown() {
	err := os.RemoveAll(TestIndexName + "/")
	if err != nil {
		log.Fatal("Could not remove test index directory")
	}
}

func TestMain(m *testing.M) {
	articles := setupArticles()
	err := NewArticleIndex(articles, TestIndexName)
	if err != nil {
		log.Fatal("Could not create testing index")
	}
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")
	teardown()
	os.Exit(exitVal)

}

func TestReadDataset(t *testing.T) {
	name := "News_Category_Dataset_v2.json"
	_, err := ReadDataset(name)
	if err != nil {
		t.Errorf("Could not read from %s correctly", name)
	}
}

func TestNewArticleIndexExists(t *testing.T) {
	// We create the test article index in TestMain
	ok := dirExists(TestIndexName)
	if !ok {
		t.Fatal("Test Article Index Does Not Exist")
	}

}

// TODO: #11 @cdpierse Fix open testing index during tests issue, may be bug
func TestMatchQuery(t *testing.T) {
	// articles := setupArticles()
	// err := NewArticleIndex(articles, TestIndexName)
	// if err != nil {
	// 	teardown()
	// 	t.Fatal("Could not create test index.")
	// }
	// // index, err := bleve.Open(TestIndexName)

	_, err := GetIndex(TestIndexName)
	if err != nil {
		t.Errorf("Could not get the test index")
	}
	// queryString := "Testing search results"
	// _, err = MatchQuery(queryString, index)
	// if err != nil {
	// 	t.Fatal("Could not create query")
	// }

}
