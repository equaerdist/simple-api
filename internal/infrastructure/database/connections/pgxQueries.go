package connections

import (
	"context"

	"github.com/equaerdist/simple-api/internal/domain/models/connections"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxQueries struct {
	Pool *pgxpool.Pool
	tx pgx.Tx
}

func NewQueries() (*PgxQueries) {
	return &PgxQueries{
		Pool: GetDBPool(),
	}
}

func(q *PgxQueries) Query(ctx context.Context,
	 query string, dest ...any) (connections.Rows, error) {
		if q.tx != nil {
			rows, err := q.tx.Query(ctx, query, dest)
			return rows, err
	 }
	return q.Pool.Query(ctx, query, dest...)
	
}

func(q *PgxQueries) QueryRow(ctx context.Context, query string,
	 dest ...any) (connections.Row) {
		if q.tx != nil {
			row := q.tx.QueryRow(ctx, query, dest)
			return row
	 }

	return q.Pool.QueryRow(ctx, query, dest...)
}

func(q *PgxQueries) WithTx(ctx context.Context) (connections.Transaction, error) {
	if q.tx  != nil {
		return q.tx, nil
	}
	return q.Pool.Begin(ctx)
}
