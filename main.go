package main

import (
	"fmt"
	"net/http"
)

func main() {
http:
	Handlefunc("/", indexHandler)
	http.HandleFunc("api/geojson", geojsonHandler)

	fs := http.FileServer
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/geojson", geoJSONHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "layout.html", nil)
}

func geoJSONHandler(w http.ResponseWriter, r *http.Request) {
	place := r.URL.Query().Get("place")
	if place == "" {
		http.Error(w, "Missing place", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("data", "geojson", strings.ToLower(place)+".geojson")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "GeoJSON not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}