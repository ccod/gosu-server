package cron

import (
	"github.com/ccod/gosu-server/pkg/config"
	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/gorilla/mux"
)

// AttachService attaches a set of routes to the router, in this case, it is blizz oauth routes
func AttachService(r *mux.Router, t *config.Tools) {
	r.HandleFunc("/cron", m.Compose(cronJob, m.PassDB(t.DB), m.PassBlizz(t.Blizz))).Methods("GET")
}
