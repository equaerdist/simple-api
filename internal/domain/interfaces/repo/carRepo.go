package repo

import (
	"context"
	"time"

	"github.com/equaerdist/simple-api/internal/domain/models"
)

type CarRepo interface {
	Get(ctx context.Context, id int) (*models.Car, error)
	Create(ctx context.Context, modelName string, createdAt time.Time) (int, error)
	Update(ctx context.Context, id int, modelName string) error
	Delete(ctx context.Context, id int) error
}
