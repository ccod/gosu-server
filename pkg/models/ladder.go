package models

// LadderRecord is only referring to 1v1
type LadderRecord struct {
	ID            int    `gorm:"AUTO_INCREMENT" json:"id"`
	Rank          int    `json:"rank"`
	League        string `json:"league"`
	Wins          int    `json:"wins"`
	Total         int    `json:"total"`
	PreferredRace string `json:"preferredRace"`
}
