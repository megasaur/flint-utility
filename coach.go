package main

import (
	"fmt"

	"github.com/go-gorp/gorp"
)

func getAllCoaches(db *gorp.DbMap) []coach {
	var coaches []coach
	db.Select(&coaches, "SELECT * FROM coach;")
	return coaches
}

func getOpponents(db *gorp.DbMap, target string) []coach {
	var coaches []coach
	// TODO: find if can do the below without the format string
	query := fmt.Sprintf("SELECT * FROM coach WHERE c_coach_id!='%s';", target)
	db.Select(&coaches, query)
	return coaches
}

func getRoundOpponent(db *gorp.DbMap, target string, round int, year int) string {
	var roundMatchupRow roundMatchup
	// TODO: find if can do the below without the format string

	query := fmt.Sprintf("SELECT * FROM round_matchup WHERE rm_round=%d AND rm_year=%d AND (rm_c_coach_id_1='%s' OR rm_c_coach_id_2='%s');", round, year, target, target)
	db.SelectOne(&roundMatchupRow, query)

	if roundMatchupRow.Coach1 == target {
		return roundMatchupRow.Coach2
	}
	return roundMatchupRow.Coach1
}
