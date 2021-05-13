package main

import (
	"log"
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
				"TMP_COUNTER",
				"ENDPOINT_URL",
			},
		}
		cgiHandler.ServeHTTP(w, r)
	}
	http.HandleFunc("/cgi-bin/", hFunc)
	if err := http.ListenAndServe("0.0.0.0:8888", nil); err != nil {
		log.Println("server error:", err)
	}

}
