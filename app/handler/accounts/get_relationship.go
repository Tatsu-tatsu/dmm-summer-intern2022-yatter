package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) GetRelationship(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	account := auth.AccountOf(r)
	username := r.URL.Query().Get("username")

	a := h.app.Dao.Account() // domain/repository の取得
	followeeAccount, err := a.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	relation := new(object.Relation)
	relation.FollowerId = account.ID
	relation.FolloweeId = followeeAccount.ID

	re := h.app.Dao.Relation()
	follow, err := re.FindRelationById(ctx, account.ID, followeeAccount.ID)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(follow); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
