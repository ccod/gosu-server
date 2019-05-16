package models

import (
	"github.com/jinzhu/gorm"
)

// History will refer to Match History, this is a straight pull from api
type History struct {
	gorm.Model
	PlayerID  int
	Map       string
	Type      string
	Decision  string
	MatchDate int // will use a date type
}

// ListHistoryByPlayer returns a full accounting of a player's match history
func ListHistoryByPlayer(db *gorm.DB, id int) []History {
	var history []History
	db.Where("player_id = ?", id).Find(&history)

	return history
}

// GetHistory returns a single Match struct given an ID
func GetHistory(db *gorm.DB, id int) History {
	var history History
	db.First(&history, id)
	return history
}

// AppendHistory saves newest additions to player's match history
func AppendHistory(db *gorm.DB, id int, newHistory []History) {
	var lastMatch History
	db.Where("player_id = ?", id).Order("match_date desc").First(&lastMatch)

	for i := 0; i < len(newHistory); i++ {
		if newHistory[i].MatchDate > lastMatch.MatchDate {
			db.Create(&newHistory[i])
		}
	}
}

// AdjudicateContest returns a boolean indicating result of match
// looking at the history of both players between the resolutionTime and 3 hours prior,
// selecting challenger history rows where both have the same match_date and type,
// taking the 3 most recent and counting successes from the challenger's point of view
func AdjudicateContest(db *gorm.DB, challengerID int, defenderID int, resolutionTime int) bool {
	var challengerHistory []History
	// TODO: fix resolutionTime interval
	db.Where(
		"match_date between ? and ? and type = 'melee' and player_id = ? and match_date in (?)", resolutionTime, resolutionTime-60*3, challengerID,
		db.Table("histories").Select("match_date").Where("match_date between ? and ? and type = 'melee' and player_id = ?", resolutionTime, resolutionTime-60*3, defenderID).QueryExpr(),
	).Order("match_date desc").Limit(3).Find(&challengerHistory)

	success := 0
	for i := 0; i < len(challengerHistory); i++ {
		if challengerHistory[i].Decision == "success" {
			success = success + 1
		}
	}

	if success == 2 {
		return true
	}

	return false
}
