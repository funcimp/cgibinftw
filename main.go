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
			InheritEnv: []string{
				"AWS_DEFAULT_REGION",
				"AWS_SESSION_TOKEN",
				"AWS_SECRET_ACCESS_KEY",
				"AWS_ACCESS_KEY_ID",
			},
		}
		cgiHandler.ServeHTTP(w, r)
	}
	http.HandleFunc("/cgi-bin/", hFunc)
	http.ListenAndServe("127.0.0.1:8888", nil)

}
