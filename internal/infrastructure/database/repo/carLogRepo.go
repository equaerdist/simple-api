package repo

import (
	"context"
	"log"
	"time"

	"github.com/equaerdist/simple-api/internal/config"
	"github.com/equaerdist/simple-api/internal/domain/interfaces/repo"
)

type CarLogRepoI struct {
	repo.Queries
}

func NewCarLogRepo(q repo.Queries) (*CarLogRepoI) {
	return &CarLogRepoI{
		Queries: q,
	}
}

func(r *CarLogRepoI) Create(ctx context.Context, carId int, modelName string,
	createdAt time.Time, updatedAt time.Time) (int, error) {
		cfg := config.GetDbCfg()
		query := "insert into " + cfg.CarLogTable + " (car_id, model_name, created_at, updated_at) values($1, $2, $3, $4) returning id"

		row := r.Queries.QueryRow(ctx, query, carId, modelName, createdAt, updatedAt)

		var logId int
		err := row.Scan(&logId)
		if err != nil {
			log.Printf("Error ocurred when try to insert car log: %v", err)
			return 0, err
		}

		return logId, nil
	}	