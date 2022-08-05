package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// POST /accounts/{username}/unfollow 機能
func (h *handler) DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	account := auth.AccountOf(r)
	username := chi.URLParam(r, "username")

	a := h.app.Dao.Account() // domain/repository の取得
	followeeAccount, err := a.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	re := h.app.Dao.Relation() // domain/repository の取得
	if err := re.DeleteRelation(ctx, account.ID, followeeAccount.ID); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// to Response body
	follow, err := re.FindRelationById(ctx, account.ID, followeeAccount.ID)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(follow); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
