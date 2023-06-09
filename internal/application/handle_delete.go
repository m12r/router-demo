package application

import "github.com/m12r/router-demo/http"

func (app *App) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		key := http.Param(r, "key")
		if err := app.db.Delete(key); err != nil {
			return err
		}
		app.respond(w, r, nil, http.StatusNoContent)
		return nil
	}
}
