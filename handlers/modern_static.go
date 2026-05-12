package handlers

import (
	"embed"
	_ "embed"
	"io/fs"
	"net/http"
)

//go:embed modern_assets_embedded/modern.html
var modernHTML []byte

//go:embed modern_assets_embedded
var modernAssetsFS embed.FS

func serveModernHTML(domainStatic string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// Manual CSP for this route
		w.Header().Set("Content-Security-Policy", "default-src 'self' data: https://gc.zgo.at https://fonts.googleapis.com https://fonts.gstatic.com https://unpkg.com; style-src 'self' 'unsafe-inline' https://gc.zgo.at https://fonts.googleapis.com; script-src 'self' 'unsafe-inline' https://gc.zgo.at https://unpkg.com; connect-src *; font-src 'self' https://fonts.gstatic.com; img-src 'self' data: https://gc.zgo.at https://fonts.gstatic.com;")
		w.Write(modernHTML)
	}
}

func serveModernAssets() http.Handler {
	f, _ := fs.Sub(modernAssetsFS, "modern_assets_embedded/modern-assets")
	return http.FileServer(http.FS(f))
}
