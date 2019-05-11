package web

import (
	"net/http"

	"github.com/vasyahuyasa/ushtr/internal/shortener"
	"github.com/vasyahuyasa/ushtr/internal/storage"
)

type HttpServer struct {
	storage   storage.GetterSaver
	shortener shortener.EncoderDecoder
	prefix    string
	mux       *http.ServeMux
}

func NewServer(urlStorage storage.GetterSaver, shortener shortener.EncoderDecoder, prefix string) *HttpServer {
	srv := &HttpServer{
		storage:   urlStorage,
		shortener: shortener,
		prefix:    prefix,
	}
	srv.makeRoutes()
	return srv
}

func (s *HttpServer) makeRoutes() {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	s.mux.Handle("/", http.HandlerFunc(s.handleSearchShortURL))
	s.mux.Handle("/-", http.HandlerFunc(s.handleCreateShortURL))
}

func (srv *HttpServer) Run(addr string) error {
	return http.ListenAndServe(addr, srv.mux)
}
