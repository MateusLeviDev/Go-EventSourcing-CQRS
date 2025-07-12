package main

import (
	"flag"
	"log"

	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/config"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/internal/server"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/logger"
)

// @contact.name Mateus Levi
// @contact.url https://github.com/MateusLeviDev
func main() {
	log.Println("Starting es microservice")

	flag.Parse()

	cfg, err := config.InitConfig()

	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("EventSourcingService")

	appLogger.Infof("CFG: %+v", cfg)
	appLogger.Info("Success =D")
	appLogger.Fatal(server.NewServer(cfg, appLogger).Run())
}
