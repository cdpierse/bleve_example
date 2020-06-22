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

	// TODO: #9 @cdpierse refactor templating
	tmpl := template.Must(template.ParseFiles("index.html"))
	handler := ServeTemplate(tmpl,articles,index)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodPost {
	// 		tmpl.Execute(w, nil)
	// 		return
	// 	}

	// 	type query struct {
	// 		queryString string
	// 	}
	// 	q := query{queryString: r.FormValue("query")}
	// 	matches, err := MatchQuery(q.queryString, index)
	// 	if err != nil {
	// 		log.Println("oops")
	// 	}
	// 	res := GetQueryHits(matches, articles)

	// 	tmpl.Execute(w, struct {
	// 		Success bool
	// 		Results Articles
	// 	}{true, res})
	// })
	http.ListenAndServe(":80", handler)

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
