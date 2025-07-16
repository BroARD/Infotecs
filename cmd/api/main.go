package main

import (
	"Infotecs/config"
	"Infotecs/internal/server"
	"Infotecs/pkg/db"
	"Infotecs/pkg/logging"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetConfig()
	logger, _ := logging.NewLogger("logs", logrus.InfoLevel)

	psqlDB, err := db.InitDB(logger)
	if err != nil {
		logger.Info("Could not start a DB")
		logger.Fatal(err)
	}

	s := server.NewServer(cfg, psqlDB, *logger)
	if err = s.Run(); err != nil {
		logger.Fatal(err)
	}
}
