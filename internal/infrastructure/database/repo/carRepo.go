package repo

import (
	"context"
	"log"
	"time"

	"github.com/equaerdist/simple-api/internal/config"
	"github.com/equaerdist/simple-api/internal/domain/interfaces/repo"
	"github.com/equaerdist/simple-api/internal/domain/models"
)

type CarRepoI struct {
	repo.Queries
}

func NewCar(q repo.Queries) (*CarRepoI) {
	return &CarRepoI{
		Queries: q,
	}
}

func(r *CarRepoI) Get(ctx context.Context, id int) (*models.Car, error) {
	cfg := config.GetDbCfg()
	query := "select * from " + cfg.CarTable + " where id = $1"
	car := models.Car{}

	row := r.Queries.QueryRow(ctx, query, id)

	err := row.Scan(&car.ID, &car.ModelName, &car.CreatedAt)
	if err != nil {
		log.Printf("error ocurred when scan data from row: %v", err)
		return nil, err
	}

	return &car, nil
}

func(r *CarRepoI) Create(ctx context.Context, modelName string,
	 createdAt time.Time) (int, error) {
	cfg := config.GetDbCfg()
	query := "insert into " + cfg.CarTable + " (model_name, created_at) values ($1, $2) returning id"

	row := r.Queries.QueryRow(ctx, query, modelName, createdAt)

	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Printf("error ocurred when try to insert model car: %v", err)
		return 0, err
	}
	
	return id, nil
}