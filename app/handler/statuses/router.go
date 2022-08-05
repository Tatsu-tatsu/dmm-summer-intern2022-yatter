package statuses

import (
	"net/http"

	"yatter-backend-go/app/app"

	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi"
)

// Implementation of handler
type handler struct {
	app *app.App
}

func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	h := &handler{app: app}
	r.With(auth.Middleware(app)).Post("/", h.Create)
	r.Get("/{id}", h.Get)

	return r
}
