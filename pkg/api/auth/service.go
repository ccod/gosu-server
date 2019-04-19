package auth

import (
	"github.com/ccod/gosu-server/pkg/config"
	"github.com/gorilla/mux"
)

// AttachService attaches a set of routes to the router, in this case, it is blizz oauth routes
func AttachService(r *mux.Router, t *config.Tools) {
	r.HandleFunc("/auth/login", login(t.AuthClient, t.OauthSalt))
	r.HandleFunc("/auth/bnet_oauth_cb", bnetCB(t.AuthClient, t.JWTSecret, t.JWTSecret))
}
