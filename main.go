package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed templates/index.html
var templateFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {
	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}
	log.Fatal(http.ListenAndServe(addr, newMux()))
}

func newMux() http.Handler {
	index := template.Must(template.ParseFS(templateFS, "templates/index.html"))

	static, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("ok\n"))
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := index.Execute(w, nil); err != nil {
			http.Error(w, "template error", http.StatusInternalServerError)
		}
	})
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(static))))
	return mux
}
