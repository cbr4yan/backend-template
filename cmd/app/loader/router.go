package loader

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func provideRouter() http.Handler {
	r := chi.NewRouter()
	return r
}
