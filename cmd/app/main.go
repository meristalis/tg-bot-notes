package main

import (
	"log"

	"github.com/meristalis/tg-bot-notes/config"
	"github.com/meristalis/tg-bot-notes/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	//new
	// Run
	//new stroke conflict
	//super new
	//conflict)
	app.Run(cfg)
}
