package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	FindStatusByID(ctx context.Context, id int64) (*object.Status, error)
	AddStatus(ctx context.Context, status object.Status) error
}
