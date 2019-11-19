package gcf

import (
	"fmt"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	v := r.URL.Query()
	if v == nil {
		return
	}
	for key, vs := range v {
		fmt.Fprintf(w, "%s = %s\n", key, vs[0])
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
