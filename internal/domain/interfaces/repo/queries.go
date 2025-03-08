package repo

import (
	"context"
	"github.com/equaerdist/simple-api/internal/domain/models/connections"
)

type Queries interface {
	Query(ctx context.Context, query string, dest ...any) (connections.Rows, error)
	QueryRow(ctx context.Context, query string, dest ...any) (connections.Row, error)
	Exec(ctx context.Context, query string, dest ...any) error
}

type Transactions interface {
	WithTx(ctx context.Context) (connections.Transaction, error)
}
