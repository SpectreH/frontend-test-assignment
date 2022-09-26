package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// SetRoutes sets handlers and load server files
func SetRoutes() http.Handler {
	mux := chi.NewRouter()

	fs := http.FileServer(http.Dir("../frontend/dist"))
	mux.Handle("/*", http.StripPrefix("/", fs))

	return mux
}
