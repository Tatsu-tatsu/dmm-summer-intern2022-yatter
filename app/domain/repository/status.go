package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// // Fetch account which has specified username
	// FindByUsername(ctx context.Context, username string) (*object.Account, error)
	// TODO: Add Other APIs
	FindStatusByID(ctx context.Context, id int64) (*object.Status, error)
	AddStatus(ctx context.Context, status object.Status) error
}
