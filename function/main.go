package main

import (
	"net/http"

	"github.com/ShintaNakama/gosolr/gcf"
)

func main() {
	http.HandleFunc("/search", gcf.Search)
	http.ListenAndServe(":8080", nil)
}
