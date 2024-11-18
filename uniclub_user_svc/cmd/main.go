package main

import (
	"log"

	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/config"
	"github.com/MohdAhzan/Uniclub_Microservice/USER_SVC/pkg/di"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./.user.env")
	if err != nil {
		log.Fatal("Error loang the env file \n", err)
	}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Couldnt load config:", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server:", diErr)
	} else {

    log.Println("user-svc running on port :7001")
		server.Start()

	}

}
