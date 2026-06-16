package main

import (
	"log"

	"github.com/00limited/football-api/internal/config"
	"github.com/00limited/football-api/internal/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := config.OpenDatabase(cfg)
	if err != nil {
		log.Fatalf("open database: %v", err)
	}

	r := router.New(cfg, db)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("run server: %v", err)
	}
}
