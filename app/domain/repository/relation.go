package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Relation interface {
	AddRelation(ctx context.Context, relation object.Relation) error
	FindRelationById(ctx context.Context, follower_id int64, followee_id int64) (*object.Follow, error)
}
