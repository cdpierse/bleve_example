package main

import (
	"bufio"
	"encoding/json"
	"github.com/blevesearch/bleve"
	"log"
	"os"
)

// Article represents a story object from the HuffPost news category dataset.
type Article struct {
	Category         string `json:"category"`
	Headline         string `json:"headline"`
	Authors          string `json:"authors"`
	Link             string `json:"link"`
	ShortDescription string `json:"short_description"`
	Date             string `json:"date"`
}

// ReadDataset reads the news dataset provided by the string `name`
// and returns a slice of `Articles`
func ReadDataset(name string) ([]Article, error) {
	data, err := os.Open(name)
	if err != nil {
		log.Panic(err)
	}
	defer data.Close()

	var articles []Article
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		var tempArticle Article
		err := json.Unmarshal(scanner.Bytes(), &tempArticle)
		if err != nil {
			return nil, err
		}
		articles = append(articles, tempArticle)
	}
	return articles, nil

}

func NewIndex(a []Article) {
	// mapping := bleve.NewIndexMapping()
	// index, err := bleve.New("example.bleve", mapping)
	index, err := bleve.Open("example.bleve")
	log.Println("Loaded index")
	if err != nil {
		panic(err)
	}

	// for _, article := range a {
	// 	index.Index(article.ShortDescription, article)
	// }

	query := bleve.NewQueryStringQuery("President Trump tweeted")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, _ := index.Search(searchRequest)
	log.Println(searchResult)

}
