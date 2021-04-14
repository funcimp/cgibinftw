package main

import (
	"net/http"
	"net/http/cgi"
)

func main() {
	hFunc := func(w http.ResponseWriter, r *http.Request) {
		cgiHandler := cgi.Handler{
			Path: "." + r.URL.Path,
			Dir:  "./dist",
		}
		cgiHandler.ServeHTTP(w, r)
	}
	http.HandleFunc("/cgi-bin/", hFunc)
	http.ListenAndServe("127.0.0.1:8888", nil)

}
