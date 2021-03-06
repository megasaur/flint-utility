package main

import (
	"fmt"
)

// WIN constant representing a won match
const WIN = "WIN"

// LOSS constant representing a lost match
const LOSS = "LOSS"

// DRAW constant representing a drawn match
const DRAW = "DRAW"

// TODO: refactor the idea of a match with these outcomes etc into a class

func getWLDTally(coach string, round int, year int) (int, int, int) {
	wins, losses, draws := 0, 0, 0
	opponents := getOpponents(coach, year)
	for _, opponent := range opponents {
		outcome, _, _ := getMatchOutcome(coach, opponent.ID, round, year)
		if outcome == WIN {
			wins++
		} else if outcome == LOSS {
			losses++
		} else {
			draws++
		}
	}
	return wins, losses, draws
}

func getRoundMatchups(round int, year int) []roundMatchup {
	// TODO: MISSING UNIT TEST
	var matchups []roundMatchup
	db := initDb()
	query := fmt.Sprintf("SELECT * FROM round_matchup where rm_year=%d and rm_round<=%d;", year, round)
	db.Select(&matchups, query)
	defer db.Db.Close()
	return matchups
}

func getMatchOutcome(coach string, opponent string, round int, year int) (string, int, int) {
	coachScores := getSaltyScores(coach, round, year)
	opponentScores := getSaltyScores(opponent, round, year)
	scoreSettings := getScoreSettings(year)
	coachTally, opponentTally := calculateScoreTallies(coachScores, opponentScores, scoreSettings)
	return getWinLossDrawStatus(coachTally, opponentTally), coachTally, opponentTally
}

func calculateScoreTallies(coachScores SaltyScores, opponentScores SaltyScores, scoreSettings scoreSettings) (int, int) {
	coachScoreTally := 0
	opponentScoreTally := 0

	if coachScores.Score > opponentScores.Score {
		coachScoreTally += scoreSettings.Score
	} else if coachScores.Score < opponentScores.Score {
		opponentScoreTally += scoreSettings.Score
	}

	if coachScores.Kicks > opponentScores.Kicks {
		coachScoreTally += scoreSettings.Kicks
	} else if coachScores.Kicks < opponentScores.Kicks {
		opponentScoreTally += scoreSettings.Kicks
	}

	if coachScores.Handballs > opponentScores.Handballs {
		coachScoreTally += scoreSettings.Handballs
	} else if coachScores.Handballs < opponentScores.Handballs {
		opponentScoreTally += scoreSettings.Handballs
	}

	if coachScores.Marks > opponentScores.Marks {
		coachScoreTally += scoreSettings.Marks
	} else if coachScores.Marks < opponentScores.Marks {
		opponentScoreTally += scoreSettings.Marks
	}

	if coachScores.Hitouts > opponentScores.Hitouts {
		coachScoreTally += scoreSettings.Hitouts
	} else if coachScores.Hitouts < opponentScores.Hitouts {
		opponentScoreTally += scoreSettings.Hitouts
	}

	if coachScores.Tackles > opponentScores.Tackles {
		coachScoreTally += scoreSettings.Tackles
	} else if coachScores.Tackles < opponentScores.Tackles {
		opponentScoreTally += scoreSettings.Tackles
	}

	if coachScores.DisposalEfficiency > opponentScores.DisposalEfficiency {
		coachScoreTally += scoreSettings.DisposalEfficiency
	} else if coachScores.DisposalEfficiency < opponentScores.DisposalEfficiency {
		opponentScoreTally += scoreSettings.DisposalEfficiency
	}

	if coachScores.ContestedPosessions > opponentScores.ContestedPosessions {
		coachScoreTally += scoreSettings.ContestedPosessions
	} else if coachScores.ContestedPosessions < opponentScores.ContestedPosessions {
		opponentScoreTally += scoreSettings.ContestedPosessions
	}

	if coachScores.Rebounds > opponentScores.Rebounds {
		coachScoreTally += scoreSettings.Rebounds
	} else if coachScores.Rebounds < opponentScores.Rebounds {
		opponentScoreTally += scoreSettings.Rebounds
	}

	if coachScores.Clearances > opponentScores.Clearances {
		coachScoreTally += scoreSettings.Clearances
	} else if coachScores.Clearances < opponentScores.Clearances {
		opponentScoreTally += scoreSettings.Clearances
	}

	if coachScores.Spoils > opponentScores.Spoils {
		coachScoreTally += scoreSettings.Spoils
	} else if coachScores.Spoils < opponentScores.Spoils {
		opponentScoreTally += scoreSettings.Spoils
	}

	return coachScoreTally, opponentScoreTally
}

func getWinLossDrawStatus(coachScoreTally int, opponentScoreTally int) string {
	if coachScoreTally > opponentScoreTally {
		return WIN
	} else if coachScoreTally < opponentScoreTally {
		return LOSS
	} else {
		return DRAW
	}
}
