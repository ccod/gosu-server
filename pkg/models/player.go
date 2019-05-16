package models

import (
	"github.com/jinzhu/gorm"
)

// Player is going to be basically my user struct
type Player struct {
	AccountID   int    `gorm:"primary_key" json:"accountId"`
	ProfileID   int    `json:"profileId"`
	ProfileURL  string `json:"profileUrl"`
	AvatarURL   string `json:"avatarUrl"`
	ClanTag     string `json:"clanTag"`
	ClanName    string `json:"clanName"`
	DisplayName string `json:"displayName"`
	RealmID     int    `json:"realmId"`
	RegionID    int    `json:"regionId"`
	Admin       bool   `gorm:"default:false" json:"admin"`

	BeefRank      int    `json:"beefRank"`
	League        string `json:"league"`
	Wins          int    `json:"wins"`
	Total         int    `json:"total"`
	PreferredRace string `json:"preferredRace"`
}

// ListPlayers returns a list of all players, will eventually do paginate
func ListPlayers(db *gorm.DB) []Player {
	var players []Player
	db.Find(&players)

	return players
}

// ListRankingPlayers returns only a set of players that have registered for Settle the Beef
func ListRankingPlayers(db *gorm.DB) []Player {
	var players []Player
	db.Where("beef_rank > 0").Find(&players)

	return players
}

// GetPlayer returns a single player given an ID
func GetPlayer(db *gorm.DB, id int) Player {
	var player Player
	db.First(&player, id)

	return player
}

// PromotePlayer moves player to designated rank, and displaces the lower ranked players down one rank
func PromotePlayer(db *gorm.DB, id int, rank int) []Player {
	var player1 Player
	var players []Player

	db.First(&player1, id)
	if player1.BeefRank == 0 {
		db.Where("rank >= ?", rank).Find(&players)
	} else {
		db.Where("rank >= ? and rank < ?", rank, player1.BeefRank).Find(&players)
	}

	player1.BeefRank = rank
	db.Save(&player1)

	for i := 0; i < len(players); i++ {
		players[i].BeefRank = players[1].BeefRank - 1
		db.Save(&players[i])
	}

	players = append(players, player1)

	return players
}

// ReplacePlayer moves player to designated rank, and unregisters the player at that rank (if applicable).
func ReplacePlayer(db *gorm.DB, id int, rank int) []Player {
	var player1, player2 Player

	db.First(&player1, id)
	db.Where("rank = ?", rank).First(&player2)

	player1.BeefRank = rank
	db.Save(&player1)

	var players = []Player{player1}

	if player2.AccountID != 0 {
		player2.BeefRank = 0
		db.Save(&player2)
		players = append(players, player2)
	}

	return players
}

// SetAdmin returns a player after saving changes to the db
func SetAdmin(db *gorm.DB, id int, val bool) Player {
	var player Player
	db.First(&player, id)
	player.Admin = val
	db.Save(&player)

	return player
}

// UnregisterPlayer returns a set of players adjusted for a player that leaves from the rankings
func UnregisterPlayer(db *gorm.DB, id int) []Player {
	var (
		player  Player
		players []Player
	)

	db.First(&player, id)
	db.Where("beef_rank > ?", player.BeefRank).Find(&players)

	player.BeefRank = 0
	db.Save(&player)

	for i := 0; i < len(players); i++ {
		players[i].BeefRank = players[i].BeefRank - 1
		db.Save(&players[i])
	}

	return append(players, player)
}

// RegisterPlayer returns player with a BeefRank at the end of the rankings
func RegisterPlayer(db *gorm.DB, id int) Player {
	var player, lastPlayer Player
	db.Order("beef_rank desc").First(&lastPlayer)
	db.First(&player, id)

	player.BeefRank = lastPlayer.BeefRank + 1
	db.Save(&player)

	return player
}

// UpdatePlayerStats updates relevant attributes from blizzard api
func UpdatePlayerStats(db *gorm.DB, oldPlayer Player, newPlayer Player) Player {
	oldPlayer.DisplayName = newPlayer.DisplayName
	oldPlayer.ClanName = newPlayer.ClanName
	oldPlayer.ClanTag = newPlayer.ClanTag
	oldPlayer.League = newPlayer.League
	oldPlayer.PreferredRace = newPlayer.PreferredRace
	oldPlayer.Wins = newPlayer.Wins
	oldPlayer.Total = newPlayer.Total
	oldPlayer.AvatarURL = newPlayer.AvatarURL

	db.Save(&oldPlayer)
	return oldPlayer
}
