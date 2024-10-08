package web

import (
	"fmt"
	"net/http"
)

func WebServer() {
	mux := http.NewServeMux()
	// files := http.FileServer(http.Dir("/piblic"))
	// mux.Handler("/static/", http.StripPrefix("/static/", files))
	fs := http.FileServer(http.Dir("/home/alex/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "URL.Path = %s", r.URL.Path)
}
