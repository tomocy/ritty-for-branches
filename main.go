package main

import (
	"log"

	"github.com/tomocy/ritty-for-branches/config"
	"github.com/tomocy/ritty-for-branches/infra/http/registerer"
	"github.com/tomocy/ritty-for-branches/infra/http/route"
	"github.com/tomocy/ritty-for-branches/infra/http/server"
	"github.com/tomocy/ritty-for-branches/registry"
)

func main() {
	config.Must(config.Load("./config.yml"))

	route.MapRoutes()

	addr := ":" + config.Current.Self.Port
	registry := registry.NewHTTPRegistry()
	apiRegi := registerer.NewAPIRegisterer(registry.NewAPIHandler())
	webRegi := registerer.NewWebRegisterer(registry.NewWebHandler())
	server := server.New(apiRegi, webRegi)
	if err := server.ListenAndServe(addr); err != nil {
		log.Printf("failed to listen and serve: %s\n", err)
	}
}
