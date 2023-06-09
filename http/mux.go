package http

import (
	"github.com/go-chi/chi/v5"
)

type Mux struct {
	mux          chi.Router
	errorHandler ErrorHandler
}

func NewMux(errorHandler ErrorHandler) *Mux {
	m := chi.NewMux()
	m.NotFound(func(w ResponseWriter, r *Request) {
		errorHandler.HandleError(w, r, &Error{StatusCode: StatusNotFound, Message: "not found"})
	})
	m.MethodNotAllowed(func(w ResponseWriter, r *Request) {
		errorHandler.HandleError(w, r, &Error{StatusCode: StatusMethodNotAllowed, Message: "method not allowed"})
	})
	return &Mux{
		mux:          m,
		errorHandler: errorHandler,
	}
}

func (m *Mux) Get(pattern string, h HandlerFunc) {
	m.Handle(MethodGet, pattern, h)
}

func (m *Mux) Post(pattern string, h HandlerFunc) {
	m.Handle(MethodPost, pattern, h)
}

func (m *Mux) Put(pattern string, h HandlerFunc) {
	m.Handle(MethodPut, pattern, h)
}

func (m *Mux) Delete(pattern string, h HandlerFunc) {
	m.Handle(MethodDelete, pattern, h)
}

func (m *Mux) Mount(pattern string, h StdHandler) {
	m.mux.Mount(pattern, h)
}

func (m *Mux) Handle(method string, pattern string, h Handler) {
	stdHandler := HandlerFuncToStd(m.errorHandler, h)
	m.mux.Method(method, pattern, stdHandler)
}

func (m *Mux) Use(middlewares ...Middleware) {
	m.mux.Use(middlewares...)
}

func (m *Mux) With(middlewares ...Middleware) *Mux {
	mux := m.mux.With(middlewares...)
	return &Mux{
		mux:          mux,
		errorHandler: m.errorHandler,
	}
}

func (m *Mux) ServeHTTP(w ResponseWriter, r *Request) {
	m.mux.ServeHTTP(w, r)
}

func HandlerFuncToStd(eh ErrorHandler, h Handler) StdHandler {
	return StdHandlerFunc(func(w ResponseWriter, r *Request) {
		if err := h.ServeHTTP(w, r); err != nil {
			eh.HandleError(w, r, err)
		}
	})
}
