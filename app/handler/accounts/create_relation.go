package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for `POST /v1/accounts`
func (h *handler) CreateRelation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	account := auth.AccountOf(r)
	username := chi.URLParam(r, "username")
	fmt.Printf(username)

	a := h.app.Dao.Account() // domain/repository の取得
	followeeAccount, err := a.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	relation := new(object.Relation)
	relation.FollowerId = account.ID
	relation.FolloweeId = followeeAccount.ID

	re := h.app.Dao.Relation() // domain/repository の取得
	if err := re.AddRelation(ctx, *relation); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	follow, err := re.FindRelationById(ctx, account.ID, followeeAccount.ID)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(follow); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
