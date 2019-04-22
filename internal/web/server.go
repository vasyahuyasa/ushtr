package web

import (
	"net/http"

	"github.com/vasyahuyasa/ushtr/internal/shortener"
	"github.com/vasyahuyasa/ushtr/internal/storage"
)

type HttpServer struct {
	storage   storage.URL
	shortener shortener.EncoderDecoder
	mux       *http.ServeMux
}

func NewServer(urlStorage storage.URL, shortener shortener.EncoderDecoder) *HttpServer {
	srv := &HttpServer{}
	srv.makeRoutes()
	return srv
}

func (s *HttpServer) makeRoutes() {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	s.mux.Handle("/", http.HandlerFunc(s.handleRequest))
}

func (srv *HttpServer) Run() error {
	return http.ListenAndServe(":9090", srv.mux)
}
