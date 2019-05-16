package history

import (
	"github.com/ccod/gosu-server/pkg/config"
	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/gorilla/mux"
)

// AttachService is the same as the other ones
func AttachService(r *mux.Router, t *config.Tools) {
	r.HandleFunc("/history/{id}", m.Compose(getHistory, m.PassDB(t.DB))).Methods("GET")
}
