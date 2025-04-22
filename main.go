package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	port := "3001"

	// Define the API endpoint handler
	http.HandleFunc("/api/geojson/", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Only allow GET requests
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract the place name from the URL
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 3 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		placeName := strings.ToLower(pathParts[len(pathParts)-1])
		filePath := filepath.Join("data", "geojson", fmt.Sprintf("%s.geojson", placeName))

		// Read the GeoJSON file
		data, err := os.ReadFile(filePath)
		if err != nil {
			// Return a 404 if the file doesn't exist
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "GeoJSON not found"})
			return
		}

		// Set content type and return the JSON data
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	// Start the server
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
