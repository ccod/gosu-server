package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ccod/gosu-server/pkg/api/player"
	"github.com/ccod/gosu-server/pkg/api/todo"

	"github.com/ccod/gosu-server/pkg/config"
	"github.com/gorilla/mux"
)

// Start pulls together services and starts server based on passed settings
func Start(c config.Settings) {
	r := mux.NewRouter()

	// smoke test, will remove later
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	}).Methods("GET")

	todo.AttachService(r, &c.Tools)
	player.AttachService(r, &c.Tools)

	fmt.Printf("Listenin on port: %s\n", c.Port)
	log.Fatal(http.ListenAndServe(":"+c.Port, r))
}
