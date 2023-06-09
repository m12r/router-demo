package application

import (
	"errors"

	"github.com/m12r/router-demo/http"
	"github.com/m12r/router-demo/internal/db"
)

func (app *App) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		key := http.Param(r, "key")
		data, err := app.db.Get(key)
		if errors.Is(err, db.ErrNotFound) {
			return &http.Error{
				StatusCode: http.StatusNotFound,
				Message:    "not found",
			}
		}
		if err != nil {
			return err
		}

		responseData := struct {
			OK   bool `json:"ok"`
			Data any  `json:"data"`
		}{
			OK:   true,
			Data: data,
		}

		app.respond(w, r, responseData, http.StatusOK)
		return nil
	}
}
