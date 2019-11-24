package gcf

import (
	"fmt"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// POSTの場合、Query()は使えない
	v := r.URL.Query()
	if v == nil {
		return
	}
	// fmt.Fprintf(w, "%t\n", v)
	host := "http://localhost:8983/solr/"
	c := "group"
	a := "select"
	q := "*:*"
	var fq string
	st := "0"
	rows := "1"
	wt := "json"
	indent := "true"
	var sort string
	var fl string
	// var recommend string

	fmt.Fprintf(w, "%s\n", r.FormValue("core"))
	fmt.Fprintf(w, "%s\n", r.FormValue("q"))
	// q = r.FormValue("q")

	fmt.Fprintf(w, "%s\n", v["q"])
	// q = v["q"]
	// fmt.Fprintf(w, "%t\n", v["q"][0])
	// fmt.Fprintf(w, "%s\n", v["q"][1])

	fmt.Println(host, c, a, q, fq, st, rows, wt, indent, sort, fl)

}
