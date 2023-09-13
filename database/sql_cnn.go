package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type SqlServerRepository struct {
	db *gorm.DB
}

func GetLocalConnection(dsn string) (*SqlServerRepository, error) {

	db, err := gorm.Open("mssql", dsn)

	if err != nil {
		return nil, err

	}

	return &SqlServerRepository{db}, nil

}
