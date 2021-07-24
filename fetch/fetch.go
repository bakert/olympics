package fetch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bakert/olympics/database"
	"github.com/jmoiron/sqlx"
)

// BAKERT move to config?
var medalTableURL = "https://og2020-api.sports.gracenote.com/svc/games_v2.svc/json/GetMedalTable_Season?competitionSetId=1&season=2021&languageCode=2"

type RawMedals struct {
	MedalTableInfo MedalTableInfo       `json:"MedalTableInfo"`
	MedalTableNOC  []MedalTableNOCEntry `json:"MedalTableNOC"`
}

type MedalTableInfo struct {
	CAsOfDate        string  `json:"c_AsOfDate"`
	NEventsTotal     int     `json:"n_EventsTotal"`
	NEventsFinished  int     `json:"n_EventsFinished"`
	NEventsScheduled int     `json:"n_EventsScheduled"`
	NMedalsGold      int     `json:"n_MedalsGold"`
	NMedalsSilver    int     `json:"n_MedalsSilver"`
	NMedalsBronze    int     `json:"n_MedalsBronze"`
	NMedalsTotal     int     `json:"n_MedalsTotal"`
	NSportID         int     `json:"n_SportID"`
	CSport           *string `json:"c_Sport"`
	CSportShort      *string `json:"c_SportShort"`
}

type MedalTableNOCEntry struct {
	NNOCID         int    `json:"n_NOCID"`
	NNOCGeoID      int    `json:"n_NOCGeoID"`
	CNOC           string `json:"c_NOC"`
	CNOCShort      string `json:"c_NOCShort"`
	NGold          int    `json:"n_Gold"`
	NSilver        int    `json:"n_Silver"`
	NBronze        int    `json:"n_Bronze"`
	NTotal         int    `json:"n_Total"`
	NRankGold      int    `json:"n_RankGold"`
	NRankSortGold  int    `json:"n_RankSortGold"`
	NRankTotal     int    `json:"n_RankTotal"`
	NRankSortTotal int    `json:"n_RankSortTotal"`
}

func UpdateMedalTable(db *sqlx.DB) error {
	b, err := fetch()
	if err != nil {
		return err
	}
	countries, err := database.Countries(db)
	if err != nil {
		return err
	}
	medals, err := parse(b, countries)
	if err != nil {
		return err
	}
	return database.Update(db, medals)
}

func fetch() ([]byte, error) {
	resp, err := http.Get(medalTableURL)
	// handle the error if there is one
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return b, err
	}
	return b, nil
}

func parse(b []byte, countries map[string]database.Country) ([]database.Medals, error) {
	var raw RawMedals
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return []database.Medals{}, err
	}
	var medals []database.Medals
	for _, r := range raw.MedalTableNOC {
		country, ok := countries[r.CNOCShort]
		if !ok {
			print("Unrecognized country: " + r.CNOC + " (" + r.CNOCShort + ")") // BAKERT handle don't drop?
			continue
		}
		medals = append(medals, database.Medals{
			Country: country,
			Medals:  r.NGold,
		})
	}
	return medals, nil
}
