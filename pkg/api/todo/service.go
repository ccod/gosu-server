package todo

import (
	"github.com/ccod/gosu-server/pkg/config"
	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/gorilla/mux"
)

// AttachService attaches a set of routes to the router
func AttachService(r *mux.Router, t *config.Tools) {
	r.HandleFunc("/todos", m.Compose(list, m.PassDB(t.DB))).Methods("GET")
	r.HandleFunc("/todos/{id}", m.Compose(index, m.PassDB(t.DB))).Methods("GET")
	r.HandleFunc("/todos", m.Compose(create, m.PassDB(t.DB))).Methods("POST")
	r.HandleFunc("/todos/{id}", m.Compose(update, m.PassDB(t.DB))).Methods("PUT")
	r.HandleFunc("/todos/{id}", m.Compose(delete, m.PassDB(t.DB))).Methods("DELETE")
	r.HandleFunc("/todos/{id}/completed", m.Compose(completed, m.PassDB(t.DB))).Methods("PUT")
}
