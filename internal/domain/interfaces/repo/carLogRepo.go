package repo

import (
	"context"
	"time"
)

type CarLogRepo interface {
	Create(ctx context.Context, carId int, modelName string,
		 createdAt time.Time, updatedAt time.Time) (int, error)
}