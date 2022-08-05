package accounts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

func (h *handler) GetFollowers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)

	a := h.app.Dao.Account() // domain/repository の取得
	followeeAccount, err := a.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	re := h.app.Dao.Relation()
	allFollowers, err := re.GetAllFollowersById(ctx, followeeAccount.ID, limit)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allFollowers); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
