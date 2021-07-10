package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed build
var fe embed.FS

func main() {
	fsys, err := fs.Sub(fe, "build")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
