package web

import (
	"log"
	"net/http"

	"github.com/vasyahuyasa/ushtr/internal/shortener"
	"github.com/vasyahuyasa/ushtr/internal/storage"
)

type shortResponse struct {
	shorturl string
}

func createShortUrl(shtnr shortener.Interface, storage storage.Interface) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := ""
		if r.Method == http.MethodGet {
			url = r.URL.Query().Get("url")
		} else if r.Method == http.MethodPost || r.Method == http.MethodPut {
			url = r.FormValue("url")
		}

		if url == "" {
			log.Println("url is empty")
			w.WriteHeader(http.StatusBadRequest)
			w.Write("Url is empty")
			return
		}

		ok, err := storage.Has(url)
		if err != nil {
			log.Printf("storage.Has(%q): %v", url, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write("Internal error")
			return
		}
	}
}
