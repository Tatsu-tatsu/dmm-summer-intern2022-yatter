package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for `POST /v1/accounts`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	statusId, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	s := h.app.Dao.Status() // domain/repository の取得
	status, err := s.FindStatusByID(ctx, statusId)
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
