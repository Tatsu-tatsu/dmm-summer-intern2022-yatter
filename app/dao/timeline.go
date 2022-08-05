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
	// Implementation for repository.Timeline
	timeline struct {
		db *sqlx.DB
	}
)

// Create timeline repository
func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func (r *timeline) FindPublicTimelines(ctx context.Context) ([]*object.Status, error) {
	entity := make([]*object.Status, 0)
	rows, err := r.db.QueryContext(ctx, "select * from status")
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
