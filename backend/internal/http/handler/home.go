package handler

import (
	"backend/internal/http/middleware"
	"backend/internal/util"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	resp := r.Context().Value(middleware.ResponseKey).(*util.Response)

	resp.Data = map[string]any{
		"message": "Welcome to API",
	}
}
