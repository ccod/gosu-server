package player

import (
	"github.com/ccod/gosu-server/pkg/config"
	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/gorilla/mux"
)

// AttachService is the same as the other ones
func AttachService(r *mux.Router, t *config.Tools) {
	// needs to be set ahead of (get player) because of the pattern matcher {id}
	r.HandleFunc("/player/register", m.Compose(mock("register player"), m.PassDB(t.DB))).Methods("GET")

	r.HandleFunc("/player", m.Compose(mock("list players"), m.PassDB(t.DB))).Methods("GET")
	r.HandleFunc("/player/{id}", m.Compose(mock("get player"), m.PassDB(t.DB))).Methods("GET")

	// I probably won't keep these three routes
	// r.HandlerFunc("/player/{id}", m.Compose(foo, m.PassDB(t.DB))).Methods("PUT")
	// r.HandlerFunc("/player/{id}", m.Compose(foo, m.PassDB(t.DB))).Methods("DELETE")
	// r.HandlerFunc("/player", m.Compose(mock("create player"), m.PassDB(t.DB))).Methods("POST")
}
