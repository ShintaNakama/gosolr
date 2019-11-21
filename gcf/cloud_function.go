package gcf

import (
	"fmt"
	"net/http"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	v := r.URL.Query()
	if v == nil {
		return
	}
	host := "http://localhost:8983/solr/"
	c := "group"
	a := "select"
	q := []string{"*:*"}
	fq := []string{}
	st := "0"
	rows := "1"
	wt := "json"
	indent := "true"
	var sort string
	var fl string
	// var recommend string

	for key, vs := range v {
		fmt.Fprintf(w, "%s = %s\n", key, vs)
		switch {
		case strings.HasPrefix(key, "core"):
			c = vs[0]
		case strings.HasPrefix(key, "q"):
			vv := key + ":"
			q = append(q, vv)
		case strings.HasPrefix(key, "fq"):
		case strings.HasPrefix(key, "action"):
			a = vs[0]
		case strings.HasPrefix(key, "start"):
			st = vs[0]
		case strings.HasPrefix(key, "rows"):
			rows = vs[0]
		case strings.HasPrefix(key, "sort"):
			sort = vs[0]
		case strings.HasPrefix(key, "fl"):

		}
	}
	// b, err := ioutil.ReadAll(r.Body)
	// fmt.Println(b)
	// if err != nil {
	// 	log.Printf("ReadAllError: %v\n", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
	// parsed, err := url.ParseQuery(string(b))
	// if err != nil {
	// 	log.Printf("ParceQueryError: %v\n", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
	// fmt.Println(toJSON(parsed))

}
