package main

import (
	"account-service/config"
	"account-service/domain"
	"account-service/grpc"
	"account-service/usecase"
	"fmt"
	"log"
)

// @contact.name Nafisa Alfiani
// @contact.email nafisa.alfiani.ica@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// init config
	cfg, err := config.InitEnv()
	if err != nil {
		log.Fatalln(err)
	}

	// init logger
	logger, err := config.InitLogger(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	logger.Info(fmt.Sprintf("%#v", cfg))

	// init DB connection
	db, err := config.InitNoSql(cfg)
	if err != nil {
		logger.Fatalf("failed to connect to mongo. %v", err)
	}

	// init domain
	dom := domain.Init(db, logger)

	// init handler
	uc := usecase.Init(cfg, logger, dom)

	g := grpc.Init(cfg, logger, uc)
	g.Run()
}
