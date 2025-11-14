package main

import (
	"log"
	"sso-service/internal/config"
	myLog "sso-service/internal/lib/log"
)

func main() {
	cfg := config.MustLoad()
	log.Println("config was loaded successfully")

	logger := myLog.MustLoad(cfg.Env)
	logger.Debug("logger was loaded successfully")
}
