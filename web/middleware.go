package web

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

func UUIDFromURLParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		threadID := chi.URLParam(r, "id")
		uuid, err := uuid.Parse(threadID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Помещаем UUID в контекст запроса
		ctx := context.WithValue(r.Context(), "threadID", uuid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
