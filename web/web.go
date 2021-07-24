package web

import (
	"fmt"
	"net/http"

	"github.com/bakert/olympics/fetch"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/bakert/olympics/database"
	"github.com/bakert/olympics/view"
)

type vars map[string]string
type handler func(http.ResponseWriter, *http.Request)
type controller func(vars, *sqlx.DB) (string, error)

func Init(db *sqlx.DB) {
	port := ":2021" // BAKERT move into config probably
	log.Info().Msg("Starting webserver on port " + port)
	r := mux.NewRouter()
	r.HandleFunc("/", makeHandler(home, db))       // BAKERT this needs to also handle 404 some way other than sending you to home
	r.HandleFunc("/{name}", makeHandler(home, db)) // BAKERT this needs to also handle 404 some way other than sending you to home
	//r.HandleFunc("/country/{id}", makeHandler(country, db))
	//r.HandleFunc("/player/{id}", makeHandler(player, db))
	//r.HandleFunc("/event/{id}", makeHandler(event, db))
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Info().Err(err).Send()
	}
	log.Info().Msg("Shutdown")
}

func makeHandler(controller controller, db *sqlx.DB) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		s, err := controller(vars, db)
		if err != nil {
			log.Error().Err(err).Send()
		}
		_, err = fmt.Fprintf(w, s)
		if err != nil {
			log.Error().Err(err).Send()
		}
	}
}

func home(_ vars, db *sqlx.DB) (string, error) {
	err := fetch.UpdateMedalTable(db)
	if err != nil {
		return "", err
	}
	log.Info().Msg("Rendering homepage")
	players, err := database.LoadPlayers(db)
	if err != nil {
		log.Error().Err(err).Send()
		return "", err
	}
	s, err := view.Render(view.Home, view.Vars{
		"title":   "Olympics Gold Medal Contest",
		"players": players,
	})
	if err != nil {
		log.Error().Err(err).Send()
		return s, err
	}
	return s, nil
}
