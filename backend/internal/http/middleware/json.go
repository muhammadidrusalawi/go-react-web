package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"backend/internal/util"
)

type ctxKey string

const ResponseKey ctxKey = "jsonResponse"

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ResponseKey, &util.Response{})
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

		resp := r.Context().Value(ResponseKey).(*util.Response)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp.Data)
	})
}
