package cron

import (
	"fmt"
	"net/http"

	"github.com/FuzzyStatic/blizzard"
	"github.com/ccod/gosu-server/pkg/client"
	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/ccod/gosu-server/pkg/models"
	re "github.com/ccod/gosu-server/pkg/response"
	"github.com/jinzhu/gorm"
)

func mock(s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from %s\n", s)
	}
}

func cronJob(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	players := models.ListPlayers(db)

	blizz := r.Context().Value(m.BlizzKey).(*blizzard.Client)
	blizz.TokenValidation()

	// cycle through each of the players, updating player stats and match history
	for i := 0; i < len(players); i++ {
		newPlayer, err := client.FetchPlayer(blizz, players[i].RegionID, players[i].RealmID, players[i].ProfileID)
		if err != nil {
			re.RespondError(err, w, r)
			return
		}

		models.UpdatePlayerStats(db, players[i], newPlayer)
		history, err := client.FetchPlayerHistory(blizz, players[i].RegionID, players[i].RealmID, players[i].ProfileID, players[i].AccountID)
		if err != nil {
			re.RespondError(err, w, r)
			return
		}

		models.AppendHistory(db, players[i].AccountID, history)
	}

	challenges := models.ListPendingAdjudicationChallenges(db)
	for i := 0; i < len(challenges); i++ {
		verdict := models.AdjudicateContest(db, challenges[i].Challenger, challenges[i].Defender, challenges[i].ResolutionDate)
		models.DecideChallenge(db, int(challenges[i].ID), verdict)

		if verdict {
			// possible spurious fetches going on here
			defender := models.GetPlayer(db, challenges[i].Defender)
			models.PromotePlayer(db, challenges[i].Challenger, defender.BeefRank)
		}
	}
}
