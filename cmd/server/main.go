package main

import (
	"go-project/internal/config"
	"go-project/internal/db/mysql"
	"go-project/internal/delivery/http"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	db, err := mysql.NewMySQLConnection(
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBName,
		cfg.DBPort,
	)

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err := mysql.AutoMigrate(db); err != nil {
		log.Fatalf("Auto-migrate failed: %v", err)
	}

	port := cfg.ServerPort
	r := http.NewRouter(cfg, db) // kirim db ke router
	if len(port) > 0 && port[0] != ':' {
		port = ":" + port
	}
	log.Fatal(r.Run(port))
}
