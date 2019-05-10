package web

import (
	"log"
	"net/http"
)

func (s *HttpServer) handleSearchShortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	code := r.URL.Path[1:]
	log.Println("Search:", code)
}

func (s *HttpServer) handleCreateShortURL(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("q")
	json := r.FormValue("json")
	log.Println("Create:", url, json)
}

func (s *HttpServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//url := r.URL.Path
		log.Println(r.URL.Path)

	case http.MethodPost:
	case http.MethodHead:
	default:
	}
}

func (s *HttpServer) index(w http.ResponseWriter) {

}

func (s *HttpServer) getUrl(w http.ResponseWriter, r *http.Request, shortURL string) {

}

func (s *HttpServer) postUrl(url string) {

}

/*
func (s *HttpServer) createShortUrl(shtnr hortener.EncoderDecoder, storage storage.Interface) http.HandleFunc {
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
*/
