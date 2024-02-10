package server

import (
	"magoto/internal/logger"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start(port int) {
	logger.Info("Starting magoto server on port %d", port)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/magoto", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/index.html")
	})

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/icons/favicon.png")
	})

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "assets"))
	FileServer(r, "/assets", filesDir)

	//TODO fix port config
	// addr := ":" + string(port)
	http.ListenAndServe(":80", r)
}

// TODO move to its own file
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
