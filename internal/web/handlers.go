package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vasyahuyasa/ushtr/internal/storage"
)

func (s *HttpServer) handleSearchShortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := r.URL.Path[1:]
	id, err := s.shortener.Decode(code)
	if err != nil {
		log.Printf("can not decode %q: %v", code, err)
		http.Error(w, "Can not decode short url", http.StatusInternalServerError)
		return
	}

	url, err := s.storage.Get(id)
	if err == storage.ErrNotFound {
		log.Printf("full url for requested code %q not found", code)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Printf("can not query full url for code %q: %v", code, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// peek
	if r.Method == http.MethodHead {
		_, err := w.Write([]byte(url))
		if err != nil {
			log.Printf("can not send peek response to client: %v", err)
		}
		return
	}

	// redirect
	http.Redirect(w, r, url, 301)
}

func (s *HttpServer) handleCreateShortURL(w http.ResponseWriter, r *http.Request) {
	// full url
	url := r.FormValue("url")

	// client want json
	needJSON := r.FormValue("json")
	isJSON := false
	if needJSON != "" && needJSON != "0" && needJSON != "false" {
		isJSON = true
	}

	// url must be defined
	if url == "" {
		http.Error(w, "URL can not be empty", http.StatusBadRequest)
		return
	}

	id, err := s.storage.Save(url)
	if err != nil {
		log.Println("can not save url to sorage:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	code := s.shortener.Encode(id)
	shortURL := s.prefix + code

	// plain text reposnse
	if !isJSON {
		_, err := fmt.Fprint(w, shortURL)
		if err != nil {
			log.Println("can not send response to client:", err)
		}
		return
	}

	// json response
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err = enc.Encode(jsonResponse{
		URL: shortURL,
	})
	if err != nil {
		log.Println("can not send JSON response to client:", err)
	}
}
