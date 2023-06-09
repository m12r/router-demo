package http

import (
	"github.com/go-chi/chi/v5"
)

func Param(r *Request, key string) string {
	return chi.URLParam(r, key)
}
