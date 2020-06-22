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

// TODO #7 : @cdpierse convert Date field to Date type.

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
func NewArticleIndex(a Articles, name string) error {
	mapping := bleve.NewIndexMapping()
	if name == "" {
		name = "example1.bleve"
	}
	index, err := bleve.New(name, mapping)
	defer index.Close()
	log.Println("Creating index")
	if err != nil {
		return err
	}

	for i, article := range a {
		index.Index(article.ID, article)
		if i%1000 == 0 {
			log.Printf("Indexed %v articles", i)
		}

	}
	return nil
}

// MatchQuery creates a new bleve MatchQuery from query string `qs`
// and returns a pointer to the search result.
func MatchQuery(qs string, index bleve.Index) (*bleve.SearchResult, error) {
	query := bleve.NewMatchQuery(qs)
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Size = 100
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

// PhraseQuery creates a new bleve PhraseQuery from a slice of term strings. The
// terms are converted to lowercase to match how they were indexed.
// Returns a pointer to the bleve.SearchResult
func PhraseQuery(terms []string, field string, index bleve.Index) (*bleve.SearchResult, error) {
	var lowerCaseTerms []string
	for _, term := range terms {
		lowerCaseTerms = append(lowerCaseTerms, strings.ToLower(term))
	}
	query := bleve.NewPhraseQuery(lowerCaseTerms, field)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	log.Printf("Search took %v seconds ", searchResult.Took)

	return searchResult, err
}

// PhraseMatchQuery ...
func PhraseMatchQuery(termPhrase string, index bleve.Index) (*bleve.SearchResult, error) {
	query := bleve.NewMatchPhraseQuery(termPhrase)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	log.Printf("Search took %v seconds ", searchResult.Took)

	return searchResult, err
}

// GetQueryHits ...
func GetQueryHits(res *bleve.SearchResult, a Articles) Articles {
	var results Articles
	hits := res.Hits
	for i := range hits {
		for j := range a {
			if hits[i].ID == a[j].ID {
				results = append(results, a[j])
				log.Printf("\nHeadline: %s\nAuthors: %s\nShort Description: %s\nDate: %s\n\n",
					a[j].Headline,
					a[j].Authors,
					a[j].ShortDescription,
					a[j].Date)
			}

		}
	}
	return results

}

//GetIndex returns a belve index of `name`
func GetIndex(name string) (bleve.Index, error) {
	index, err := bleve.Open(name)
	if err != nil {
		return nil, err
	}
	return index, nil

}
