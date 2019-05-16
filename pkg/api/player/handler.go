package player

import (
	"fmt"
	"net/http"
	"strconv"

	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/ccod/gosu-server/pkg/models"
	re "github.com/ccod/gosu-server/pkg/response"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func mock(s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from %s\n", s)
	}
}

// admin actions
func addAdmin(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	player := models.SetAdmin(db, id, true)
	re.RespondJSON(player, w, r)
}

func removeAdmin(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	player := models.SetAdmin(db, id, false)
	re.RespondJSON(player, w, r)
}

func adminPlayerRegister(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	players := models.UnregisterPlayer(db, id)
	re.RespondJSON(players, w, r)
}

func adminPlayerUnregister(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	players := models.UnregisterPlayer(db, id)
	re.RespondJSON(players, w, r)
}

func replacePlayerRank(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	rank, err := strconv.Atoi(mux.Vars(r)["rank"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	players := models.ReplacePlayer(db, id, rank)
	re.RespondJSON(players, w, r)
}

func promotePlayerRank(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	rank, err := strconv.Atoi(mux.Vars(r)["rank"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	players := models.PromotePlayer(db, id, rank)
	re.RespondJSON(players, w, r)
}

// user actions
func listPlayers(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)

	players := models.ListPlayers(db)
	re.RespondJSON(players, w, r)
}

func listRankingPlayers(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	players := models.ListRankingPlayers(db)
	re.RespondJSON(players, w, r)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	player := models.GetPlayer(db, id)
	re.RespondJSON(player, w, r)
}

func listPlayerChallenges(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	challenges := models.ListChallengesByPlayer(db, id)
	re.RespondJSON(challenges, w, r)
}

func listPlayerHistory(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	history := models.ListHistoryByPlayer(db, id)
	re.RespondJSON(&history, w, r)
}

func registerPlayer(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	accID := r.Context().Value(m.JWTKey).(int)

	player := models.RegisterPlayer(db, accID)
	re.RespondJSON(player, w, r)
}

func unregisterPlayer(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	accID := r.Context().Value(m.JWTKey).(int)

	players := models.UnregisterPlayer(db, accID)
	re.RespondJSON(players, w, r)
}
