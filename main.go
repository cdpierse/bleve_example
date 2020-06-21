package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

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
		NewArticleIndex(articles, "")
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
	// TODO: #9 @cdpierse refactor templating
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		type query struct {
			queryString string
		}
		q := query{queryString: r.FormValue("query")}
		matches, err := MatchQuery(q.queryString, index)
		if err != nil {
			log.Println("oops")
		}
		res := GetQueryHits(matches, articles)

		tmpl.Execute(w, struct {
			Success bool
			Results Articles
		}{true, res})
	})
	http.ListenAndServe(":80", nil)

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
