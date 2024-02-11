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

var assetDir string

func Start(port int) {
	logger.Info("Starting magoto server on port %d", port)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//TODO embed assets or store in config dir
	//use config dir if dynamic compilation is needed
	exDir, err := os.Executable()
	if err != nil {
		logger.FatalErr("Error getting working directory: %s", err)
	}
	assetDir = filepath.Join(filepath.Dir(exDir), "assets")

	r.Get("/magoto", func(w http.ResponseWriter, r *http.Request) {
		fpath := asset("index.html")
		logger.Debug("Serving file: %s", fpath)
		http.ServeFile(w, r, fpath)
	})

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, asset("icons/favicon.png"))
	})

	FileServer(r, "/assets", http.Dir(assetDir))

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
		logger.Debug("Path prefix: %s", pathPrefix)
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func asset(relPath string) string {
	return filepath.Join(assetDir, relPath)
}
