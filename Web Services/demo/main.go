package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("www")))
	google, _ := url.Parse("http://www.google.com")
	http.Handle("/google", httputil.NewSingleHostReverseProxy(google))

	http.HandleFunc("/mock", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", `{"name":"Sean Murphy"}`)
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
