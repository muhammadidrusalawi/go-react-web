package http

import (
	"net/http"

	"backend/internal/http/handler"
	"backend/internal/http/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/product", middleware.JSONMiddleware(http.HandlerFunc(handler.Product)))
	mux.Handle("/home", middleware.JSONMiddleware(http.HandlerFunc(handler.Home)))

	return middleware.LoggingMiddleware(mux)
}
