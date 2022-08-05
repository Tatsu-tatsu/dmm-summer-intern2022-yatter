package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	relation struct {
		db *sqlx.DB
	}
)

// Create relation repository
func NewRelation(db *sqlx.DB) repository.Relation {
	return &relation{db: db}
}

func (r *relation) AddRelation(ctx context.Context, relation object.Relation) error {
	// relationship := new(object.Relation)
	_, err := r.db.ExecContext(ctx, "INSERT INTO relation (follower_id, followee_id) VALUES (?, ?)", relation.FollowerId, relation.FolloweeId)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (r *relation) FindRelationById(ctx context.Context, follower_id int64, followee_id int64) (*object.Follow, error) {
	follow := new(object.Follow)
	relationship := new(object.Relation)
	err := r.db.QueryRowxContext(ctx, "select * from relation where follower_id = ? and followee_id = ?", follower_id, followee_id).StructScan(relationship)
	if err == nil {
		follow.Following = true
	} else if errors.Is(err, sql.ErrNoRows) {
		follow.Following = false
	} else {
		return nil, fmt.Errorf("%w", err)
	}

	relationshipReverse := new(object.Relation)
	err = r.db.QueryRowxContext(ctx, "select * from relation where follower_id = ? and followee_id = ?", followee_id, follower_id).StructScan(relationshipReverse)
	if err == nil {
		follow.Followed_by = true
	} else if errors.Is(err, sql.ErrNoRows) {
		follow.Followed_by = false
	} else {
		return nil, fmt.Errorf("%w", err)
	}

	return follow, nil
}

func (r *relation) DeleteRelation(ctx context.Context, follower_id int64, followee_id int64) error {
	_, err := r.db.ExecContext(ctx, "delete from relation where follower_id = ? and followee_id = ?", follower_id, followee_id)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
