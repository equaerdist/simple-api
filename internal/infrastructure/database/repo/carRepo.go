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

func NewCar(q repo.Queries) *CarRepoI {
	return &CarRepoI{
		Queries: q,
	}
}

func (r *CarRepoI) Get(ctx context.Context, id int) (*models.Car, error) {
	cfg := config.GetDbCfg()
	query := "select * from " + cfg.CarTable + " where id = $1"
	car := models.Car{}

	row, err := r.Queries.QueryRow(ctx, query, id)

	if err != nil {
		log.Printf("Error ocurred when exec row: %v", err)
		return nil, err
	}

	err = row.Scan(&car.ID, &car.ModelName, &car.CreatedAt)
	if err != nil {
		log.Printf("error ocurred when scan data from row: %v", err)
		return nil, err
	}

	return &car, nil
}

func (r *CarRepoI) Create(ctx context.Context, modelName string,
	createdAt time.Time) (int, error) {
	cfg := config.GetDbCfg()
	query := "insert into " + cfg.CarTable + " (model_name, created_at) values ($1, $2) returning id"

	row, err := r.Queries.QueryRow(ctx, query, modelName, createdAt)
	if err != nil {
		log.Printf("Error ocurred when exec row: %v", err)
		return 0, err
	}

	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Printf("error ocurred when try to insert model car: %v", err)
		return 0, err
	}

	return id, nil
}

func (r *CarRepoI) Update(ctx context.Context, id int, modelName string) error {
	cfg := config.GetDbCfg()
	query := "update " + cfg.CarTable + " set model_name = $1 where id = $2"

	err := r.Queries.Exec(ctx, query, modelName, id)
	if err != nil {
		log.Printf("Error ocurred when try to update car row: %v", err)
		return err
	}

	return nil
}

func (r *CarRepoI) Delete(ctx context.Context, id int) error {
	cfg := config.GetDbCfg()
	query := "delete from " + cfg.CarTable + " where id = $1"

	err := r.Queries.Exec(ctx, query, id)
	if err != nil {
		log.Printf("Error ocurred when try to delete car row: %v", err)
		return err
	}

	return nil
}
