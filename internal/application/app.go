package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"strconv"

	"github.com/m12r/router-demo/http"
	"github.com/m12r/router-demo/internal/db"
)

type App struct {
	mux *http.Mux
	db  db.DB
}

func NewApp(db db.DB) *App {
	app := &App{
		db: db,
	}
	app.mux = http.NewMux(app.errorHandler())
	app.routes()
	return app
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.mux.ServeHTTP(w, r)
}

func (app *App) respond(w http.ResponseWriter, r *http.Request, data any, statusCode int) {
	if data == nil {
		w.WriteHeader(statusCode)
		return
	}

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		log.Printf("cannot encode to json: %v", err)
		buf.Reset()
		buf.WriteString(`{"ok": false, code": 500, "message": "internal server error"}`)
	}

	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = io.Copy(w, buf)
}

func (app *App) errorHandler() http.ErrorHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		var httpErr *http.Error
		if !errors.As(err, &httpErr) {
			httpErr = &http.Error{
				StatusCode: 500,
				Message:    "internal server error",
				Err:        err,
			}
		}

		responseData := struct {
			OK      bool   `json:"ok"`
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			OK:      false,
			Code:    httpErr.StatusCode,
			Message: httpErr.Message,
		}

		app.respond(w, r, responseData, httpErr.StatusCode)
	}
}
