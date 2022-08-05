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
	// Implementation for repository.Relation
	relation struct {
		db *sqlx.DB
	}
)

// Create relation repository
func NewRelation(db *sqlx.DB) repository.Relation {
	return &relation{db: db}
}

func (r *relation) AddRelation(ctx context.Context, relation object.Relation) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO relation (follower_id, followee_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE follower_id = follower_id", relation.FollowerId, relation.FolloweeId)

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

	// フォローする人とフォローされる人を反対にして検索
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

func (r *relation) GetAllFollowingsById(ctx context.Context, follower_id int64, limit int64) ([]*object.Account, error) {
	entity := make([]*object.Account, 0)
	rows, err := r.db.QueryContext(ctx, "select account.* from relation INNER JOIN account ON account.id = relation.followee_id where follower_id = ? LIMIT ?", follower_id, limit)
	err = sqlx.StructScan(rows, &entity)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	defer rows.Close()

	return entity, nil
}

func (r *relation) GetAllFollowersById(ctx context.Context, followee_id int64, limit int64, since_id int64, max_id int64) ([]*object.Account, error) {
	entity := make([]*object.Account, 0)
	var rows *sql.Rows
	var err error
	// max_idに記入がないとき、初期値0となるので条件から排除
	if max_id == 0 {
		rows, err = r.db.QueryContext(ctx, "select account.* from relation INNER JOIN account ON account.id = relation.follower_id where followee_id = ? and ? <= follower_id LIMIT ?", followee_id, since_id, limit)
	} else {
		rows, err = r.db.QueryContext(ctx, "select account.* from relation INNER JOIN account ON account.id = relation.follower_id where followee_id = ? and ? <= follower_id and follower_id <= ? LIMIT ?", followee_id, since_id, max_id, limit)
	}
	err = sqlx.StructScan(rows, &entity)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	defer rows.Close()

	return entity, nil
}

func (r *relation) DeleteRelation(ctx context.Context, follower_id int64, followee_id int64) error {
	_, err := r.db.ExecContext(ctx, "delete from relation where follower_id = ? and followee_id = ?", follower_id, followee_id)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
