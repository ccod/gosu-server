package main

import (
	"github.com/ccod/gosu-server/pkg/api"
	"github.com/ccod/gosu-server/pkg/config"
)

func main() {
	// will directly panic if it isn't able to find necessary vars
	settings := config.LoadFromEnv()
	defer settings.Tools.DB.Close()

	api.Start(settings)
}
