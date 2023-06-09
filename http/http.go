package http

import (
	"net/http"
)

type ResponseWriter = http.ResponseWriter

type Request = http.Request

type StdHandler = http.Handler

type StdHandlerFunc = http.HandlerFunc

type Server = http.Server

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request) error
}

type HandlerFunc func(w ResponseWriter, r *Request) error

func (fn HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) error {
	return fn(w, r)
}
