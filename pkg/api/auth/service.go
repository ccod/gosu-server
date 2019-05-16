package auth

import (
	"github.com/ccod/gosu-server/pkg/config"
	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/gorilla/mux"
)

// AttachService attaches a set of routes to the router, in this case, it is blizz oauth routes
func AttachService(r *mux.Router, t *config.Tools) {
	r.HandleFunc("/auth/login", login(t.AuthClient, t.OauthSalt)).Methods("GET")
	r.HandleFunc("/auth/bnet_oauth_cb", bnetCB(t.AuthClient, t.JWTSecret))
	r.HandleFunc("/auth/user", m.Compose(getUser, m.JWTIdentity(t.JWTSecret), m.PassDB(t.DB), m.PassBlizz(t.Blizz))).Methods("GET")
}
