package models

import (
	"time"
)

type Car struct {
	ID         int
	ModelName string
	CreatedAt time.Time
}

type CarLog struct {
	ID int
	LogId int
	ModelName string
	CreatedAt time.Time
}