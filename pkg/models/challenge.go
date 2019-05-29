package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Challenge is the crux of the purpose of this application, with respect to internal clan ladder
// TODO: think about having a place holder for Player struct for challenger and defender
type Challenge struct {
	gorm.Model
	Challenger     int       `json:"challenger"`
	Defender       int       `json:"defender"`
	Decision       string    `gorm:"default:undecided" json:"decision"` // prefer an enum [undecided, success, failure] with respect to Challenger
	ResolutionDate time.Time `json:"resolutionDate"`                    // TODO: Make this time.Time and comparable to History.MatchDate
}

// ListChallengesByPlayer returns a list of challenges as both challenger and defender
func ListChallengesByPlayer(db *gorm.DB, id int) []Challenge {
	var challenges []Challenge
	db.Where("challenger = ? or defender = ?", id, id).Find(&challenges)

	return challenges
}

// ListChallenges returns a full list of all the challenges... not likely to be used
func ListChallenges(db *gorm.DB) []Challenge {
	var challenges []Challenge
	db.Find(&challenges)

	return challenges
}

// ListUnresolvedChallenges returns all current challenges with discision being undecided
func ListUnresolvedChallenges(db *gorm.DB) []Challenge {
	var challenges []Challenge
	db.Where("decision = 'undecided'").Find(&challenges)

	return challenges
}

// ListPendingAdjudicationChallenges returns a list of challenges awaiting adjudication
func ListPendingAdjudicationChallenges(db *gorm.DB) []Challenge {
	var challenges []Challenge
	db.Where("decision = 'undecided' and resolution_date is not null").Find(&challenges)

	return challenges
}

// GetChallenge returns a single challenge given an id
func GetChallenge(db *gorm.DB, id int) Challenge {
	var challenge Challenge
	db.First(&challenge, id)

	return challenge
}

// CreateChallenge returns a new challenge
func CreateChallenge(db *gorm.DB, challenger int, defender int) Challenge {
	challenge := Challenge{Challenger: challenger, Defender: defender}
	db.Create(&challenge)
	return challenge
}

// DeleteChallenge will soft delete the challenge
func DeleteChallenge(db *gorm.DB, id int) Challenge {
	var challenge Challenge
	db.First(&challenge, id)
	db.Delete(&challenge)

	return challenge
}

// DecideChallenge returns an edited challenge that decided a challenger victor
func DecideChallenge(db *gorm.DB, id int, verdict bool) Challenge {
	var challenge Challenge
	db.First(&challenge, id)
	if verdict == true {
		challenge.Decision = "success"
	} else {
		challenge.Decision = "failure"
	}
	db.Save(&challenge)

	return challenge
}

// QueueChallengeResolution sets a timestamp at the end of the contest indicating roughly when the competition was settled
// need to change issueDate, and resolutionDate to Time.time and validate string
func QueueChallengeResolution(db *gorm.DB, id int, t int64) Challenge {
	var challenge Challenge
	db.First(&challenge, id)
	challenge.ResolutionDate = time.Unix(t, 0)
	db.Save(&challenge)

	return challenge
}
