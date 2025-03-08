package config

import (
	"sync"
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
		dbCfgInstance = &DbCfg{
			Dsn: "postgresql://myuser:mypassword@localhost:6543/mydb",
			CarTable: "car",
			CarLogTable: "car_log",
		}
	})
	return dbCfgInstance
}
