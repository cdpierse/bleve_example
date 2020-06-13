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

// NewArticleIndex buils a new bleve index from the
// `Articles` passed.
// TODO: #2 #1 Add a name parameter
func NewArticleIndex(a Articles) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example1.bleve", mapping)
	log.Println("Creating index")
	if err != nil {
		panic(err)
	}

	for i, article := range a {
		index.Index(article.ID, article)
		if i%1000 == 0 {
			log.Printf("Indexed %v articles", i)
		}

	}
}

// MatchQueryIndex ...
func MatchQueryIndex(qs string, index bleve.Index) (*bleve.SearchResult, error) {
	query := bleve.NewMatchQuery(qs)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	log.Printf("Search took %v seconds ", searchResult.Took)

	return searchResult, err
}

// GetQueryHits ...
func GetQueryHits(res *bleve.SearchResult, a Articles) {
	hits := res.Hits
	for i := range hits {
		for j := range a {
			if hits[i].ID == a[j].ID {
				log.Printf("\nHeadline: %s\nAuthors: %s\nShort Description: %s\nDate: %s\n\n",
					a[j].Headline,
					a[j].Authors,
					a[j].ShortDescription,
					a[j].Date)
			}

		}
	}

}

//GetIndex returns a belve index of `name`
func GetIndex(name string) (bleve.Index, error) {
	return bleve.Open(name)
}
