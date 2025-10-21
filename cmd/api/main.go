package main

import (
	"log"

	"github.com/ItsMeSouptik/insightly/internal/config"
	"github.com/ItsMeSouptik/insightly/internal/db"
	"github.com/ItsMeSouptik/insightly/internal/models"
	"github.com/ItsMeSouptik/insightly/internal/server"
)

func main() {
	cfg := config.Load()
	db.Connect()
	err := db.DB.AutoMigrate(&models.Note{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("✅ Database migrated!")
	server.Run(cfg)
}
