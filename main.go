package main

import (
	"flag"

	"github.com/bakert/olympics/database"
	"github.com/bakert/olympics/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

type olympicsConfig struct {
	dsn string
}

func main() {
	log.Info().Msg("Olympics system startup")
	cfg := config()
	db, err := database.Init(cfg.dsn)
	if err != nil {
		panic(err)
	}
	web.Init(db)
}

// BAKERT this doesn't actually work
func config() olympicsConfig {
	log.Info().Msg("Configuring")
	cfg := olympicsConfig{}
	flag.StringVar(&cfg.dsn, "dsn", "olympics:olympics@/olympics", "DSN for mariadb database")
	flag.Parse()
	return cfg
}
