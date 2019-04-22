package web

import (
	"net/http"

	"github.com/vasyahuyasa/ushtr/internal/shortener"
	"github.com/vasyahuyasa/ushtr/internal/storage"
)

type HttpServer struct {
	storage *storage.Interface
	shrt    *shortener.Interface
	mux     *http.ServeMux
}

func (s *HttpServer) makeRoutes() {

}
