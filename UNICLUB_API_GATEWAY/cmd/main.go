package main

import (
	"log"

	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/config"
	"github.com/MohdAhzan/Uniclub_Microservice/API_GATEWAY/pkg/di"
)

func main() {

	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		log.Fatal("cannot load config: ", cfgErr)
	}

	server, diErr := di.InitializeAPI(cfg)

	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
