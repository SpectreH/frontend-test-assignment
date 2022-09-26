package handlers

import (
	"net/http"
	"test-assgnment/internal/routes/config"
)

// Index is the handler for vue routing
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/favicon.ico" {
		http.ServeFile(w, r, config.UI_DIST+r.URL.Path)
		return
	}

	http.ServeFile(w, r, config.UI_DIST+"/index.html")
}
