package accounts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

func (h *handler) GetFollowing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")
	limitQuery := r.URL.Query().Get("limit")
	limit, _ := strconv.ParseInt(limitQuery, 10, 64)

	const defaultLimit int64 = 40
	if limitQuery == "" {
		limit = defaultLimit
	}

	a := h.app.Dao.Account() // domain/repository の取得
	followingAccount, err := a.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// {username}のフォローをすべて取得
	re := h.app.Dao.Relation()
	allFollowings, err := re.GetAllFollowingsById(ctx, followingAccount.ID, limit)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(allFollowings); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
