package web

import (
	"embed"
	"net/http"
)

var (
	// go:embed web
	StaticWeb embed.FS
)

func DoWebStatic(w http.ResponseWriter, req *http.Request) {

}
