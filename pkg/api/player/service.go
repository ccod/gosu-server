package player

import (
	"github.com/ccod/gosu-server/pkg/config"
	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/gorilla/mux"
)

// AttachService is the same as the other ones
func AttachService(r *mux.Router, t *config.Tools) {
	r.HandleFunc("/player", m.Compose(listPlayers, m.PassDB(t.DB))).Methods("GET")
	r.HandleFunc("/player/rankings", m.Compose(listRankingPlayers, m.PassDB(t.DB))).Methods("GET")

	// uses jwt claim for these routes
	r.HandleFunc("/player/register", m.Compose(registerPlayer, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret), m.PassAdmin)).Methods("GET")
	r.HandleFunc("/player/unregister", m.Compose(unregisterPlayer, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret), m.PassAdmin)).Methods("GET")

	r.HandleFunc("/player/{id}", m.Compose(getPlayer, m.PassDB(t.DB))).Methods("GET")

	r.HandleFunc("/player/{id}/challenges", m.Compose(listPlayerChallenges, m.PassDB(t.DB))).Methods("GET")
	r.HandleFunc("/player/{id}/history", m.Compose(listPlayerHistory, m.PassDB(t.DB))).Methods("GET")

	r.HandleFunc("/player/{id}/remove-admin", m.Compose(adminPlayerUnregister, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret), m.PassAdmin)).Methods("PUT")
	r.HandleFunc("/player/{id}/add-admin", m.Compose(adminPlayerUnregister, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret), m.PassAdmin)).Methods("PUT")

	r.HandleFunc("/player/{id}/ranking/unregister", m.Compose(adminPlayerUnregister, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret), m.PassAdmin)).Methods("PUT")
	r.HandleFunc("/player/{id}/ranking/register", m.Compose(adminPlayerRegister, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret), m.PassAdmin)).Methods("PUT")
	r.HandleFunc("/player/{id}/replace/{rank}", m.Compose(replacePlayerRank, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret), m.PassAdmin)).Methods("PUT")
	r.HandleFunc("/player/{id}/promote/{rank}", m.Compose(promotePlayerRank, m.PassDB(t.DB), m.JWTIdentity(t.JWTSecret), m.PassAdmin)).Methods("PUT")

	// I probably won't keep these three CRUD routes
	// r.HandlerFunc("/player/{id}", m.Compose(foo, m.PassDB(t.DB))).Methods("PUT")
	// r.HandlerFunc("/player/{id}", m.Compose(foo, m.PassDB(t.DB))).Methods("DELETE")
	// r.HandlerFunc("/player", m.Compose(mock("create player"), m.PassDB(t.DB))).Methods("POST")
}
