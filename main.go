package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ts := time.Now()
		data := map[string]string{
			"CurrentTime": ts.UTC().Format("Mon Jan 2 15:04:05 MST 2006"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
		log.Println(ts, "/")
	})
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("listening on", port)
	http.ListenAndServe(":"+port, mux)
}
