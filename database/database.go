package database

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type Player struct {
	ID int
	Name string
}

func Init(dsn string) (*sqlx.DB, error) {
	log.Info().Msg("Opening database connection with " + dsn) // BAKERT remove logging of pwd
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return db, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

func LoadPlayers(db *sqlx.DB) ([]Player, error) {
	sql := `
		SELECT 
			p.id, 
			p.name 
		FROM
			player AS p`
	var players = make([]Player, 2)
	err := db.Select(&players, sql)
	if err != nil {
		return nil, err
	}
	return players, nil
}