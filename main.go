package main

import (
	"log"

	"github.com/tomocy/ritty-for-branches/config"
	"github.com/tomocy/ritty-for-branches/infra/http/registerer"
	"github.com/tomocy/ritty-for-branches/infra/http/server"
)

func main() {
	config.Must(config.Load("./config.yml"))

	addr := ":" + config.Current.Self.Port
	webRegi := registerer.NewWebRegisterer()
	server := server.New(webRegi)
	if err := server.ListenAndServe(addr); err != nil {
		log.Printf("failed to listen and serve: %s\n", err)
	}
}
