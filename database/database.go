package database

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type Country struct {
	ID     int
	Name   string
	Code   string
	Last   int
	Medals int
}

type Player struct {
	ID        int
	Name      string
	Countries []Country
	Medals    int
}

type playerRow struct {
	PlayerID      int    `db:"player_id"`
	PlayerName    string `db:"player_name"`
	CountryID     int    `db:"country_id"`
	CountryName   string `db:"country_name"`
	CountryCode   string `db:"country_code"`
	CountryLast   int    `db:"country_last"`
	CountryMedals int    `db:"country_medals"`
}

type Medals struct {
	Country Country
	Medals  int
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
			p.id AS player_id, 
			p.name AS player_name,
		    c.id AS country_id,
		    c.name AS country_name,
		    c.code AS country_code,
		    c.last AS country_last,
		    c.medals AS country_medals
		FROM
			player AS p
		LEFT JOIN
			player_country AS pc ON p.id = pc.player_id
		INNER JOIN
			country AS c ON pc.country_id = c.id
		GROUP BY
		    p.id,
			c.id,
			c.medals,
		    c.last,
		    c.name
		ORDER BY
		    p.id,
			c.medals DESC,
		    c.last DESC,
			c.name
		`
	var results []playerRow
	err := db.Select(&results, sql)
	if err != nil {
		return nil, err
	}
	var players []Player
	var player Player
	for _, r := range results {
		if r.PlayerID != player.ID {
			if player.ID > 0 {
				players = append(players, player)
				player = Player{}
			}
			player.ID = r.PlayerID
			player.Name = r.PlayerName
			player.Countries = []Country{}
		}
		country := Country{
			Name:   r.CountryName,
			Code:   r.CountryCode,
			Last:   r.CountryLast,
			Medals: r.CountryMedals,
		}
		player.Countries = append(player.Countries, country)
		player.Medals = player.Medals + country.Medals
	}
	players = append(players, player)
	return players, nil
}

func Update(db *sqlx.DB, medals []Medals) error {
	sql := "UPDATE country SET medals = ? WHERE id = ?"
	for _, m := range medals {
		_, err := db.Exec(sql, m.Medals, m.Country.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func Countries(db *sqlx.DB) (map[string]Country, error) {
	sql := "SELECT id, name, code, last, medals FROM country"
	var results []Country
	err := db.Select(&results, sql)
	if err != nil {
		return nil, err
	}
	countries := make(map[string]Country)
	for _, r := range results {
		countries[r.Code] = r
	}
	return countries, nil
}
