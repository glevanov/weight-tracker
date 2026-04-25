package static

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed dist/.keep dist/*
var distFS embed.FS

func Handler() (http.Handler, bool) {
	content, err := fs.Sub(distFS, "dist")
	if err != nil {
		return nil, false
	}

	fileServer := http.FileServer(http.FS(content))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			http.NotFound(w, r)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		if _, err := fs.Stat(content, path); err != nil {
			r.URL.Path = "/index.html"
		}

		fileServer.ServeHTTP(w, r)
	}), true
}
