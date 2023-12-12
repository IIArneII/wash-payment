package main

import (
	"log"
	"wash-payment/internal/app"
	"wash-payment/internal/config"
	"wash-payment/internal/dal"
	"wash-payment/internal/pkg/db"
	"wash-payment/internal/pkg/logger"
	"wash-payment/internal/services"
	"wash-payment/internal/services/rabbit"
	"wash-payment/internal/transport/firebase"
	rabbitHandler "wash-payment/internal/transport/rabbit"
	"wash-payment/internal/transport/rest"

	"go.uber.org/zap"
)

/*
TODO
Основные задачи

	DONE1)Добавить методы списания средств по group_id или organisation_id(service/rabbit):

	DONE2)Подтянуть Бд с ShareBuisnes(transport/rabbit/sendMessage)

	DONE3)Добавить Upserve методы в Круды (Services->Заменить Update и Create на один, который
	вызовет нужный метод из repo)

	4)	Переписать тесты для 3 пункта

	5)	Протестировать организации (получение через рэббит, обновление, что все встает в бонусную и что бонусная делает рассылку)
*/

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("init config: ", err)
	}

	l, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalln("init logger: ", err)
	}
	l.Info("logger initialized")

	dbConn, err := db.InitDB(cfg.DB.DSN())
	if err != nil {
		l.Fatalln("init db: ", err)
	}
	defer dbConn.Close()
	l.Info("connected to db")

	err = db.UpMigrations(dbConn.DB, cfg.DB.MigrationsDir)
	if err != nil {
		l.Fatalln("migrate db: ", err)
	}
	l.Info("migrations applied")

	repositories := dal.NewRepositories(l, dbConn)
	services := services.NewServices(l, repositories)

	authSvc, err := firebase.NewFirebaseService(l, cfg.FirebaseConfig.FirebaseKeyFilePath, services.UserService)
	if err != nil {
		log.Fatalln("init firebase service: ", err)
	}
	l.Info("connected firebase")

	rabbitSvc := rabbit.NewService(l, services)

	_, err = rabbitHandler.NewRabbitService(l, cfg.RabbitMQConfig, rabbitSvc)

	if err != nil {
		log.Fatalln("init rabbit service: ", err)
	}
	l.Info("connected rabbit")

	errc := make(chan error)
	go runHTTPServer(errc, l, cfg, services, authSvc)

	err = <-errc
	if err != nil {
		l.Fatalln("http server: ", err)
	}
}

func runHTTPServer(errc chan error, l *zap.SugaredLogger, cfg config.Config, services *app.Services, authSvc firebase.FirebaseService) {
	defer func() {
		if r := recover(); r != nil {
			l.Fatalln("panic: ", r)
		}
	}()

	l.Info("http server started")
	defer l.Info("http server finished")

	server, err := rest.NewServer(l, cfg, services, authSvc)
	if err != nil {
		l.Fatalln("Init http server: ", err)
	}

	errc <- server.Serve()
}
