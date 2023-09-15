package main

import (
	"log"
	"wash-payment/internal/config"
	"wash-payment/internal/pkg/db"
	"wash-payment/internal/pkg/logger"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("Init config: ", err)
	}

	l, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalln("Init logger: ", err)
	}

	dbConn, err := db.InitDB(cfg.DB)
	if err != nil {
		l.Fatalln("Init db: ", err)
	}
	defer dbConn.Close()

	err = db.UpMigrations(dbConn.DB, cfg.DB.MigrationsDir)
	if err != nil {
		l.Fatalln("Migrate db: ", err)
	}
}
