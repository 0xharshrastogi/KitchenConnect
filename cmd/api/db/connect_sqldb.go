package db

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectSqlServerDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, err
}
