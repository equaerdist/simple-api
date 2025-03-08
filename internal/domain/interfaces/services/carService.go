package services

import (
	"context"

	"github.com/equaerdist/simple-api/internal/domain/models"
)

type CarService interface {
	GetCar(ctx context.Context, id int) (*models.Car, error)
	CreateCar(ctx context.Context, modelName string) (int, error)
}