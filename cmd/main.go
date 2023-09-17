package main

import (
	"log"
	"wash-payment/internal/app"
	"wash-payment/internal/config"
	"wash-payment/internal/dal"
	"wash-payment/internal/pkg/db"
	"wash-payment/internal/pkg/logger"
	"wash-payment/internal/services"
	"wash-payment/internal/transport/firebase"
	"wash-payment/internal/transport/rest"

	"go.uber.org/zap"
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
	l.Info("Logger initialized")

	dbConn, err := db.InitDB(cfg.DB)
	if err != nil {
		l.Fatalln("Init db: ", err)
	}
	defer dbConn.Close()
	l.Info("Connected to db")

	err = db.UpMigrations(dbConn.DB, cfg.DB.MigrationsDir)
	if err != nil {
		l.Fatalln("Migrate db: ", err)
	}
	l.Info("Migrations applied")

	dal := dal.NewDal(l, dbConn)
	services := services.NewServices(l, dal)

	authSvc, err := firebase.New(cfg.FirebaseConfig.FirebaseKeyFilePath, l, services.UserService)

	errc := make(chan error)
	go runHTTPServer(errc, l, cfg, services, authSvc)

	err = <-errc
	if err != nil {
		l.Fatalln("HTTP server: ", err)
	}
}

func runHTTPServer(errc chan error, l *zap.SugaredLogger, cfg config.Config, services *app.Services, authSvc firebase.FirebaseService) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("Panic: ", r)
		}
	}()

	l.Info("HTTP server started")
	defer l.Info("HTTP server finished")

	server, err := rest.NewServer(l, cfg, services, authSvc)
	if err != nil {
		log.Fatalln("Init HTTP server: ", err)
	}

	errc <- server.Serve()
}
