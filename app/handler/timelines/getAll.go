package timelines

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	t := h.app.Dao.Timeline() // domain/repository の取得
	status, err := t.FindPublicTimelines(ctx)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

}
