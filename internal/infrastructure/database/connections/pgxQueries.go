package connections

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"

	"github.com/equaerdist/simple-api/internal/domain/models/connections"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxQueries struct {
	Pool *pgxpool.Pool
	tx   pgx.Tx
	con  *pgxpool.Conn
}

func NewQueries() *PgxQueries {
	return &PgxQueries{
		Pool: GetDBPool(),
	}
}

func (q *PgxQueries) Query(ctx context.Context,
	query string, dest ...any) (connections.Rows, error) {
	if q.tx != nil {
		rows, err := q.tx.Query(ctx, query, dest...)
		return rows, err
	}

	if q.con != nil {
		return q.con.Query(ctx, query, dest...)
	}

	_, err := q.createConnectionInternal(ctx)
	if err != nil {
		return nil, err
	}

	return q.con.Query(ctx, query, dest...)

}

func (q *PgxQueries) QueryRow(ctx context.Context, query string,
	dest ...any) (connections.Row, error) {
	if q.tx != nil {
		row := q.tx.QueryRow(ctx, query, dest...)
		return row, nil
	}

	if q.con != nil {
		return q.Pool.QueryRow(ctx, query, dest...), nil
	}

	_, err := q.createConnectionInternal(ctx)
	if err != nil {
		return nil, err
	}

	return q.Pool.QueryRow(ctx, query, dest...), nil
}

func (q *PgxQueries) WithTx(ctx context.Context) (connections.Transaction, error) {
	if q.tx != nil {
		return q.tx, nil
	}
	con, err := q.Pool.Acquire(ctx)

	if err != nil {
		log.Printf("error acquiring connections: %v", err)
		return nil, err
	}

	q.con = con
	tx, err := q.con.Begin(ctx)

	if err != nil {
		log.Printf("error beginning transaction: %v", err)
		return nil, err
	}
	q.tx = tx

	return q.tx, nil
}

func (q *PgxQueries) createConnectionInternal(ctx context.Context) (*pgxpool.Conn, error) {
	con, err := q.Pool.Acquire(ctx)
	if err != nil {
		log.Printf("error acquiring connections: %v", err)
		return nil, err
	}

	q.con = con

	return q.con, nil
}

func (q *PgxQueries) Exec(ctx context.Context, query string, dest ...any) error {
	if q.tx != nil {
		_, err := q.tx.Exec(ctx, query, dest...)
		return err
	}

	if q.con != nil {
		_, err := q.con.Exec(ctx, query, dest...)
		return err
	}

	_, err := q.createConnectionInternal(ctx)
	if err != nil {
		return err
	}

	_, err = q.con.Exec(ctx, query, dest...)
	return err
}
