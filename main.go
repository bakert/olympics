package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/bakert/olympics/view"
)

func main() {
	port := ":2021"
	log.Info().Msg("Starting webserver on port " + port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Info().Err(err).Send()
	}
	log.Info().Msg("Shutdown")
}

func handler(w http.ResponseWriter, r *http.Request) {
	s, err := view.Render(view.Home, map[string]interface{}{
		"title": "Olympics Gold Medal Contest",
		"players": players(),
	})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	_, err = fmt.Fprintf(w, s)
	if err != nil {
		log.Error().Err(err).Send()
	}
}

type Player struct {
	Name      string
	Countries []Country
}

type Country struct {
	Name      string
	Code      string
	Golds     int
	LastGolds int
}

func players() []Player {
	return []Player{
		{
			Name: "Katie",
			Countries: []Country{
				{
					Name:      "United States",
					Code:      "USA",
					Golds:     0,
					LastGolds: 46,
				},
				{
					Name:      "Germany",
					Code:      "GER",
					Golds:     0,
					LastGolds: 17,
				},
				{
					Name:      "Japan",
					Code:      "JPN",
					Golds:     0,
					LastGolds: 12,
				},
				{
					Name:      "France",
					Code:      "FRA",
					Golds:     0,
					LastGolds: 10,
				},
				{
					Name:      "South Korea",
					Code:      "KOR",
					Golds:     0,
					LastGolds: 9,
				},
				{
					Name:      "Australia",
					Code:      "AUS",
					Golds:     0,
					LastGolds: 8,
				},
				{
					Name:      "Netherlands",
					Code:      "NED",
					Golds:     0,
					LastGolds: 8,
				},
				{
					Name:      "Hungary",
					Code:      "HUN",
					Golds:     0,
					LastGolds: 8,
				},
				{
					Name:      "Spain",
					Code:      "ESP",
					Golds:     0,
					LastGolds: 7,
				},
				{
					Name:      "Kenya",
					Code:      "KEN",
					Golds:     0,
					LastGolds: 6,
				},
				{
					Name:      "Cuba",
					Code:      "CUB",
					Golds:     0,
					LastGolds: 5,
				},
				{
					Name:      "New Zealand",
					Code:      "NZL",
					Golds:     0,
					LastGolds: 4,
				},
				{
					Name:      "Argentina",
					Code:      "ARG",
					Golds:     0,
					LastGolds: 3,
				},
				{
					Name:      "Denmark",
					Code:      "DEN",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "Sweden",
					Code:      "SWE",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "South Africa",
					Code:      "RSA",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "Serbia",
					Code:      "SRB",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "Portugal",
					Code:      "POR",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Belgium",
					Code:      "BEL",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "Slovakia",
					Code:      "SVK",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "Belarus",
					Code:      "BLR",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Slovenia",
					Code:      "SLO",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Indonesia",
					Code:      "INA",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Romania",
					Code:      "ROU",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Bahrain",
					Code:      "BRN",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Vietnam",
					Code:      "VIE",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Ivory Coast",
					Code:      "CIV",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Fiji",
					Code:      "FIJ",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Puerto Rico",
					Code:      "PUR",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Singapore",
					Code:      "SIN",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Mexico",
					Code:      "MEX",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Algeria",
					Code:      "ALG",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Ireland",
					Code:      "IRL",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Bulgaria",
					Code:      "BUL",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Mongolia",
					Code:      "MGL",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Niger",
					Code:      "NIG",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Philippines",
					Code:      "PHI",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Qatar",
					Code:      "QAT",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Tunisia",
					Code:      "TUN",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Israel",
					Code:      "ISR",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Austria",
					Code:      "AUT",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Nigeria",
					Code:      "NGR",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "United Arab Emirates",
					Code:      "UAE",
					Golds:     0,
					LastGolds: 0,
				},
			},
		},
		{
			Name:      "Tom",
			Countries: []Country{
				{
					Name:      "Great Britain",
					Code:      "GBR",
					Golds:     0,
					LastGolds: 27,
				},
				{
					Name:      "China",
					Code:      "CHN",
					Golds:     0,
					LastGolds: 26,
				},
				{
					Name:      "Russia",
					Code:      "RUS",
					Golds:     0,
					LastGolds: 19,
				},
				{
					Name:      "Japan",
					Code:      "JPN",
					Golds:     0,
					LastGolds: 12,
				},
				{
					Name:      "Italy",
					Code:      "ITA",
					Golds:     0,
					LastGolds: 8,
				},
				{
					Name:      "Brazil",
					Code:      "BRA",
					Golds:     0,
					LastGolds: 7,
				},
				{
					Name:      "Jamaica",
					Code:      "JAM",
					Golds:     0,
					LastGolds: 6,
				},
				{
					Name:      "Czech Republic",
					Code:      "CZE",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Croatia",
					Code:      "CRO",
					Golds:     0,
					LastGolds: 5,
				},
				{
					Name:      "Canada",
					Code:      "CAN",
					Golds:     0,
					LastGolds: 4,
				},
				{
					Name:      "Uzbekistan",
					Code:      "UZB",
					Golds:     0,
					LastGolds: 4,
				},
				{
					Name:      "Colombia",
					Code:      "COL",
					Golds:     0,
					LastGolds: 3,
				},
				{
					Name:      "Switzerland",
					Code:      "SUI",
					Golds:     0,
					LastGolds: 3,
				},
				{
					Name:      "Kazakhstan",
					Code:      "KAZ",
					Golds:     0,
					LastGolds: 3,
				},
				{
					Name:      "Iran",
					Code:      "IRI",
					Golds:     0,
					LastGolds: 3,
				},
				{
					Name:      "Greece",
					Code:      "GRE",
					Golds:     0,
					LastGolds: 3,
				},
				{
					Name:      "Ukraine",
					Code:      "UKR",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "Poland",
					Code:      "POL",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "North Korea",
					Code:      "PRK",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "Thailand",
					Code:      "THA",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "Georgia",
					Code:      "GEO",
					Golds:     0,
					LastGolds: 2,
				},
				{
					Name:      "Azerbaijan",
					Code:      "AZE",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Turkey",
					Code:      "TUR",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Armenia",
					Code:      "ARM",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Ethiopia",
					Code:      "ETH",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Chinese Taipei",
					Code:      "TPE",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Bahamas",
					Code:      "BAH",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Independent Olympic Athletes",
					Code:      "IOA",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Jordan",
					Code:      "JOR",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Kosovo",
					Code:      "KOS",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Tajikistan",
					Code:      "TJK",
					Golds:     0,
					LastGolds: 1,
				},
				{
					Name:      "Malaysia",
					Code:      "MAS",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Venezuela",
					Code:      "VEN",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Lithuania",
					Code:      "LTU",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "India",
					Code:      "IND",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Burundi",
					Code:      "BDI",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Grenada",
					Code:      "GRN",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Norway",
					Code:      "NOR",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Egypt",
					Code:      "EGY",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Dominican Republic",
					Code:      "DOM",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Estonia",
					Code:      "EST",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Finland",
					Code:      "FIN",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Morocco",
					Code:      "MAR",
					Golds:     0,
					LastGolds: 0,
				},
				{
					Name:      "Trinidad and Tobago",
					Code:      "TTO",
					Golds:     0,
					LastGolds: 0,
				},
			},
		},
	}
}