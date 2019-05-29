package challenge

import (
	"encoding/json"
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

func listChallenges(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	challenges := models.ListChallenges(db)
	re.RespondJSON(challenges, w, r)
}

func listUnresolvedChallenges(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	challenges := models.ListUnresolvedChallenges(db)
	re.RespondJSON(challenges, w, r)
}

func getChallenge(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	challenge := models.GetChallenge(db, id)
	re.RespondJSON(challenge, w, r)
}

func createChallenge(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)

	var i map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		re.RespondError(err, w, r)
		return
	}

	challenge := models.CreateChallenge(db, i["challenger"].(int), i["defender"].(int))
	re.RespondJSON(challenge, w, r)
}

func queueChallengeResolution(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	var i map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		re.RespondError(err, w, r)
		return
	}

	challenge := models.QueueChallengeResolution(db, id, i["resolutionDate"].(int64))
	re.RespondJSON(challenge, w, r)
}

func adjudicateChallenge(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	var i map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		re.RespondError(err, w, r)
		return
	}

	challenge := models.DecideChallenge(db, id, i["descision"].(bool))
	re.RespondJSON(challenge, w, r)
}

func deleteChallenge(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		re.RespondError(err, w, r)
		return
	}

	challenge := models.DeleteChallenge(db, id)
	re.RespondJSON(challenge, w, r)
}
