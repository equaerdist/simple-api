package connections

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/equaerdist/simple-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
    dbPool *pgxpool.Pool
    once   sync.Once
)

func GetDBPool() *pgxpool.Pool {
	cfg := config.GetDbCfg()
    once.Do(func() {
        var err error
        dbPool, err = pgxpool.New(context.Background(), cfg.Dsn)
	
        if err != nil {
            log.Fatalf("Unable to connect to database: %v", err)
        }
        fmt.Println("Database connection pool established")
    })
    return dbPool
}
