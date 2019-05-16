package challenge

import (
	"github.com/ccod/gosu-server/pkg/config"
	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/gorilla/mux"
)

// AttachService attaches a set of routes to the router, in this case, it is blizz oauth routes
func AttachService(r *mux.Router, t *config.Tools) {
	r.HandleFunc("/challenge", m.Compose(listChallenges, m.PassDB(t.DB))).Methods("GET")
	r.HandleFunc("/challenge/unresolved", m.Compose(listUnresolvedChallenges, m.PassDB(t.DB))).Methods("GET")

	r.HandleFunc("/challenge/{id}", m.Compose(getChallenge, m.PassDB(t.DB))).Methods("GET")

	// if challenger is not provided in post params, infer it from token
	r.HandleFunc("/challenge", m.Compose(createChallenge, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret))).Methods("POST")

	// user ability to queue up a job for adjudication
	r.HandleFunc("/challenge/{id}/queue-resolution", m.Compose(queueChallengeResolution, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret))).Methods("PUT")
	// admin ability to decide on a challenge
	r.HandleFunc("/challenge/{id}/adjudicate", m.Compose(adjudicateChallenge, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret), m.PassAdmin)).Methods("PUT")

	// This is for revoked challenges (the challenger opts to not follow through), or denied challenges (defender has already recieved a challenge recently from challenger and opts out)
	r.HandleFunc("/challenge/{id}", m.Compose(deleteChallenge, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret))).Methods("DELETE")
}
