package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/blevesearch/bleve"
)

// Article represents a story object from the HuffPost news category dataset.
type Article struct {
	ID               string
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
	id := 0
	for scanner.Scan() {
		var tempArticle Article
		tempArticle.ID = strconv.Itoa(id)
		id++
		err := json.Unmarshal(scanner.Bytes(), &tempArticle)
		if err != nil {
			return nil, err
		}
		articles = append(articles, tempArticle)
	}
	return articles, nil

}

// Articles represents a collection of articles i.e from a database
// or some other datastore.
type Articles []Article

func NewArticleIndex(a Articles) {
	// mapping := bleve.NewIndexMapping()
	// index, err := bleve.New("example.bleve", mapping)
	index, err := bleve.Open("example.bleve")
	log.Println("Loaded index")
	if err != nil {
		panic(err)
	}

	// for i, article := range a[:5000] {
	// 	index.Index(article.ID, article)
	// 	if i%1000 == 0 {
	// 		log.Printf("Indexed %v articles", i)
	// 	}

	// }


	query := bleve.NewMatchQuery("Actors on strike syria due to war")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, _ := index.Search(searchRequest)
	log.Println(searchResult)

	for _, article := range a {
		for j := range searchResult.Hits {
			if article.ID == searchResult.Hits[j].ID {
				log.Println(article)
			}
		}
	}

}

// func QueryIndex(qs string, index *bleve.Index) bleve.SearchResult {
// 	query := bleve.NewQueryStringQuery(qs)
// 	searchRequest := bleve.NewSearchRequest(query)
// 	searchRequest.Highlight
// 	return
// }
