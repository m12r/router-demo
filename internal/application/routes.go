package application

func (app *App) routes() {
	app.mux.Get("/{key}", app.handleGet())
	app.mux.Put("/{key}", app.handlePut())
	app.mux.Delete("/{key}", app.handleDelete())
}
