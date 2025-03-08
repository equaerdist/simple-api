package config

import (
	"os"
	"sync"

	"github.com/equaerdist/simple-api/internal/domain/consts"
)

type DbCfg struct {
	Dsn string
	CarTable string
	CarLogTable string
}

var (
	dbCfgInstance *DbCfg
	once          sync.Once
)

func GetDbCfg() *DbCfg {
	once.Do(func() {
		dbDsn := os.Getenv(consts.DB_DSN)
		dbCfgInstance = &DbCfg{
			Dsn: dbDsn,
			CarTable: consts.CAR,
			CarLogTable: consts.CAR_LOG,
		}
	})
	return dbCfgInstance
}
