package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	Db *sqlx.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func (p *PostgresRepository) InitDB(cfg Config) (*sqlx.DB, error) {
	pg, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = pg.Ping()
	if err != nil {
		return nil, err
	}
	p.Db = pg
	return pg, nil
}
