package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/blevesearch/bleve"
)

type query struct {
	queryString string
}
// ServeTemplate returns a handler function for handling
// the executed tempalte and the search queries posted to it
func ServeTemplate(tpl *template.Template, articles Articles, index bleve.Index) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tpl.Execute(w, nil)
			return
		}

		q := query{queryString: r.FormValue("query")}
		matches, err := MatchQuery(q.queryString, index)
		if err != nil {
			log.Println("oops")
		}
		res := GetQueryHits(matches, articles)
		tpl.Execute(w, struct {
			Success bool
			Results Articles
		}{true, res})

	}
	return http.HandlerFunc(fn)

}
