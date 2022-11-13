package main

import (
	"company/internal/application"
	"log"
)

func main() {

	//Creating application instance
	app := application.New(
		application.Name(`company-management`),
	)

	//Initialise application
	if err := app.Init(); err != nil {
		log.Fatal(err.Error(), `action`, `application.Init`)
	}

	//Run application
	if err := app.Run(); err != nil {
		log.Fatal(err.Error(), `action`, `application.Run`)
	}
}
