package database

import (
	"fmt"

	"github.com/HermanPlay/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) Database {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=disable",
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDatabase{Db: db}
}

func (p *postgresDatabase) Connect() *gorm.DB {
	return p.Db
}
