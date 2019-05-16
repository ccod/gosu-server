package client

import (
	"strconv"

	"github.com/FuzzyStatic/blizzard"
	"github.com/ccod/gosu-server/pkg/models"
)

// FetchNewPlayer takes an account ID and blizzard Client, and returns a filled player struct
func FetchNewPlayer(blizz *blizzard.Client, i int) (models.Player, error) {
	blizz.TokenValidation()
	var player models.Player

	playerData, _, err := blizz.SC2Player(i)
	if err != nil {
		return player, err
	}

	profileID, err := strconv.Atoi((*playerData)[0].ProfileID)
	if err != nil {
		return player, err
	}

	playerRegion := blizzard.Region((*playerData)[0].RegionID)

	sc2Profile, _, err := blizz.SC2Profile(playerRegion, (*playerData)[0].RealmID, profileID)
	if err != nil {
		return player, err
	}

	sc2LegacyProfile, _, err := blizz.SC2LegacyProfile(playerRegion, (*playerData)[0].RealmID, profileID)
	if err != nil {
		return player, err
	}

	player = models.Player{
		AccountID:     i,
		ProfileID:     profileID,
		ProfileURL:    (*playerData)[0].ProfileURL,
		AvatarURL:     (*playerData)[0].AvatarURL,
		ClanTag:       sc2LegacyProfile.ClanTag,
		ClanName:      sc2LegacyProfile.ClanName,
		DisplayName:   (*playerData)[0].Name,
		RealmID:       (*playerData)[0].RealmID,
		RegionID:      (*playerData)[0].RegionID,
		BeefRank:      0,
		League:        sc2Profile.Snapshot.SeasonSnapshot.OneV1.LeagueName.(string),
		Wins:          sc2Profile.Snapshot.SeasonSnapshot.OneV1.TotalWins,
		Total:         sc2Profile.Snapshot.SeasonSnapshot.OneV1.TotalGames,
		PreferredRace: sc2LegacyProfile.Career.PrimaryRace,
	}

	return player, nil
}

// FetchPlayer mainly used to update player stats for currently existing players
func FetchPlayer(blizz *blizzard.Client, region int, realm int, profileID int) (models.Player, error) {
	var player models.Player
	playerRegion := blizzard.Region(region)

	sc2Profile, _, err := blizz.SC2Profile(playerRegion, realm, profileID)
	if err != nil {
		return player, err
	}

	sc2LegacyProfile, _, err := blizz.SC2LegacyProfile(playerRegion, realm, profileID)
	if err != nil {
		return player, err
	}

	player = models.Player{
		AvatarURL:     sc2Profile.Summary.Portrait,
		ClanTag:       sc2LegacyProfile.ClanTag,
		ClanName:      sc2LegacyProfile.ClanName,
		DisplayName:   sc2Profile.Summary.DisplayName,
		League:        sc2Profile.Snapshot.SeasonSnapshot.OneV1.LeagueName.(string),
		Wins:          sc2Profile.Snapshot.SeasonSnapshot.OneV1.TotalWins,
		Total:         sc2Profile.Snapshot.SeasonSnapshot.OneV1.TotalGames,
		PreferredRace: sc2LegacyProfile.Career.PrimaryRace,
	}

	return player, nil
}

// FetchPlayerHistory returns a collection of matches for a given player
func FetchPlayerHistory(blizz *blizzard.Client, region int, realm int, profileID int, accID int) ([]models.History, error) {
	var history []models.History
	regionID := blizzard.Region(region)
	profileMatches, _, err := blizz.SC2LegacyProfileMatches(regionID, realm, profileID)
	if err != nil {
		return history, err
	}

	for i := 0; i < len((*profileMatches).Matches); i++ {
		history = append(history, models.History{
			PlayerID:  accID,
			Map:       (*profileMatches).Matches[i].Map,
			Type:      (*profileMatches).Matches[i].Type,
			Decision:  (*profileMatches).Matches[i].Decision,
			MatchDate: (*profileMatches).Matches[i].Date,
		})
	}

	return history, nil
}
