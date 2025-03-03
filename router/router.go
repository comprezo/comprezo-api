package router

import (
	"comprezo/config"
	"comprezo/handler"
	"comprezo/router/handlers"
	"net/http"
)

func Init(cfg *config.Config) *http.ServeMux {
	r := http.NewServeMux()
	initRoutes(r)
	return r
}

func initRoutes(r *http.ServeMux) {
	r.Handle("GET /{$}", handler.REST(handlers.Home))
	r.Handle("GET /get-size", handler.REST(handlers.GetSize))
}
