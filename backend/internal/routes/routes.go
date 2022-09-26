package routes

import (
	"net/http"

	"test-assgnment/internal/handlers"
	"test-assgnment/internal/routes/config"

	"github.com/go-chi/chi"
)

// SetRoutes sets handlers and load server files
func SetRoutes() http.Handler {
	mux := chi.NewRouter()

	fs := http.FileServer(http.Dir(config.UI_DIST))
	mux.Handle("/static/*", fs)

	mux.Get("/*", handlers.Index)

	return mux
}
