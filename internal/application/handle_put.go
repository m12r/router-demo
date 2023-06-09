package application

import (
	"io"

	"github.com/m12r/router-demo/http"
)

func (app *App) handlePut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		value, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		key := http.Param(r, "key")
		if err := app.db.Put(key, string(value)); err != nil {
			return err
		}
		app.respond(w, r, nil, http.StatusNoContent)
		return nil
	}
}
