package main

import (
	"log"

	"github.com/jayghz/go-backend/cmd/app"
)

func main() {
	cfg := app.Config{
		Addr: ":4000",
	}

	application := app.Application{
		Config: cfg,
	}

	mux := application.Mount()
	log.Fatal(application.Run(mux))
}
