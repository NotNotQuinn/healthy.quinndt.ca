package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/healthy.quinndt.ca/templates"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(templates.FilePathPrefix + templates.Index)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	h := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// https://developers.google.com/search/docs/crawling-indexing/robots-meta-tag
		w.Header().Set("X-Robots-Tag", "noindex, nofollow")
		mux.ServeHTTP(w, r)
	}))
	http.ListenAndServe("127.0.0.1:3002", h) // blocks
}
