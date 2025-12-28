package handler

import (
	"net/http"

	"backend/internal/http/middleware"
	"backend/internal/util"
)

func Product(w http.ResponseWriter, r *http.Request) {
	resp := r.Context().Value(middleware.ResponseKey).(*util.Response)

	resp.Data = map[string]any{
		"message": "List of products",
		"products": []map[string]any{
			{
				"id":    1,
				"name":  "Keyboard",
				"price": 250000,
			},
			{
				"id":    2,
				"name":  "Mouse",
				"price": 150000,
			},
		},
	}
}
