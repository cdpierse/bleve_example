package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

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

// MatchQuery creates a new bleve MatchQuery from query string `qs`
// and returns a pointer to the search result.
func MatchQuery(qs string, index bleve.Index) (*bleve.SearchResult, error) {
	query := bleve.NewMatchQuery(qs)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	log.Printf("Search took %v seconds ", searchResult.Took)

	return searchResult, err
}

// TermQuery creates a new bleve TermQuery from query from the term
// `t` and returns a pointer to the search result.
func TermQuery(t string, index bleve.Index) (*bleve.SearchResult, error) {
	t = strings.ToLower(t)
	query := bleve.NewTermQuery(t)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	log.Printf("Search took %v seconds ", searchResult.Took)

	return searchResult, err
}

func PhraseQuery(terms []string, field string, index bleve.Index) (*bleve.SearchResult, error) {
	// TODO: #5 Add loop to change terms to lowercase 
	query := bleve.NewPhraseQuery(terms, field)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	log.Printf("Search took %v seconds ", searchResult.Took)

	return searchResult, err
}

func PhraseMatchQuery(termPhrase string, index bleve.Index) (*bleve.SearchResult, error) {
	query := bleve.NewMatchPhraseQuery(termPhrase)
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
