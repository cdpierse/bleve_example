package main

// import (
// 	"html/template"
// 	_ "log"
// 	"net/http"
// )

// func GetTemplate() *template.Template {
// 	tpl := template.Must(template.ParseFiles("index.html"))
// 	return tpl

// }

// func ServeTemplate() *http.Handler {
// 	tpl := GetTemplate()
// 	return http.HandleFunc("/",
// 		func(w http.ResponseWriter, r *http.Request) {
// 			tpl.Execute(w, nil)

// 		})
// }
