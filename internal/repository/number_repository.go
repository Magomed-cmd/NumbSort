package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type QueryExecutor interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

type numberRepository struct {
	db QueryExecutor
}

func NewNumberRepository(db QueryExecutor) *numberRepository {
	return &numberRepository{db: db}
}

func (r *numberRepository) Insert(ctx context.Context, value int) error {
	_, err := r.db.Exec(ctx, "INSERT INTO numbers(value) VALUES($1)", value)
	return err
}

func (r *numberRepository) ListSorted(ctx context.Context) ([]int, error) {
	rows, err := r.db.Query(ctx, "SELECT value FROM numbers ORDER BY value ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []int
	for rows.Next() {
		var v int
		if scanErr := rows.Scan(&v); scanErr != nil {
			return nil, scanErr
		}
		result = append(result, v)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
