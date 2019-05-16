package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ccod/gosu-server/pkg/api/auth"
	"github.com/ccod/gosu-server/pkg/api/challenge"
	"github.com/ccod/gosu-server/pkg/api/cron"
	"github.com/ccod/gosu-server/pkg/api/history"
	"github.com/ccod/gosu-server/pkg/api/player"

	"github.com/ccod/gosu-server/pkg/api/todo"

	"github.com/ccod/gosu-server/pkg/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Start pulls together services and starts server based on passed settings
func Start(c config.Settings) {
	r := mux.NewRouter()

	todo.AttachService(r, &c.Tools)

	auth.AttachService(r, &c.Tools)
	player.AttachService(r, &c.Tools)
	challenge.AttachService(r, &c.Tools)
	history.AttachService(r, &c.Tools)
	cron.AttachService(r, &c.Tools)

	cor := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,

		Debug: true,
	})

	fmt.Printf("Listenin on port: %s\n", c.Port)
	log.Fatal(http.ListenAndServe(":"+c.Port, cor.Handler(r)))
}
