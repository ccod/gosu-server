package models

import "github.com/jinzhu/gorm"

// Ranking refers to the internal ranking of gosu beef players in the game of "Settle the Beef"
type Ranking struct {
	gorm.Model
	Rank     int    `json:"rank"`
	Player   Player `gorm:"foreignkey:PlayerID" json:"player"`
	PlayerID int    `json:"playerId"`
}
