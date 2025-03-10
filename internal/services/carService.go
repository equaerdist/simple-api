package services

import (
	"context"
	"log"
	"time"

	"github.com/equaerdist/simple-api/internal/domain/interfaces/repo"
	"github.com/equaerdist/simple-api/internal/domain/models"
	"github.com/equaerdist/simple-api/internal/infrastructure/database/connections"
	repos "github.com/equaerdist/simple-api/internal/infrastructure/database/repo"
)

type CarServiceI struct {
	repo.Transactions
	repo.CarRepo
	repo.CarLogRepo
}

func NewCarService() *CarServiceI {
	queries := connections.NewQueries()
	carRepo := repos.NewCar(queries)
	carLogRepo := repos.NewCarLogRepo(queries)

	return &CarServiceI{
		Transactions: queries,
		CarRepo:      carRepo,
		CarLogRepo:   carLogRepo,
	}
}

func (s *CarServiceI) GetCar(ctx context.Context, id int) (*models.Car, error) {
	tx, err := s.Transactions.WithTx(ctx)
	if err != nil {
		log.Printf("Error ocurred when start tx: %v", err)
		return nil, err
	}

	car, err := s.CarRepo.Get(ctx, id)
	if err != nil {
		log.Printf("Error ocurred when try to get car from repo: %v", err)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Printf("Error ocurred when try to commit tx:  %v", err)
		return nil, err
	}

	return car, nil
}

func (s *CarServiceI) CreateCar(ctx context.Context, modelName string) (int, error) {
	tx, err := s.Transactions.WithTx(ctx)
	if err != nil {
		log.Printf("Error ocurred when start tx: %v", err)
		return 0, err
	}
	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			log.Printf("Error ocurred when rollback tx: %v", err)
		}
	}()

	createdAt := time.Now()
	id, err := s.CarRepo.Create(ctx, modelName, createdAt)
	if err != nil {
		log.Printf("Error ocurred when try to create car in service: %v", err)
		return 0, err
	}
	_, err = s.CarLogRepo.Create(ctx, id, modelName, createdAt, time.Now())
	if err != nil {
		log.Printf("Error ocurred when try to log car creation in service: %v", err)
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Printf("Error ocurred when try to commit tx:  %v", err)
		return 0, err
	}

	return id, nil
}

func (s *CarServiceI) UpdateCar(ctx context.Context, id int, modelName string) error {
	tx, err := s.Transactions.WithTx(ctx)
	if err != nil {
		log.Printf("Error ocurred when start tx: %v", err)
		return err
	}
	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			log.Printf("Error ocurred when rollback tx: %v", err)
		}
	}()

	car, err := s.CarRepo.Get(ctx, id)
	if err != nil {
		return err
	}

	err = s.CarRepo.Update(ctx, id, modelName)
	_, err = s.CarLogRepo.Create(ctx, id, modelName, car.CreatedAt, time.Now())
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	return err
}

func (s *CarServiceI) DeleteCar(ctx context.Context, id int) error {
	tx, err := s.Transactions.WithTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			log.Printf("Error ocurred when rollback tx: %v", err)
		}
	}()

	car, err := s.CarRepo.Get(ctx, id)
	if err != nil {
		return err
	}

	err = s.CarRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	_, err = s.CarLogRepo.Create(ctx, id, car.ModelName, car.CreatedAt, time.Now())
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	return err
}
