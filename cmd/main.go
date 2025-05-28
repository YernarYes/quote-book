package main

import (
	"quotes/config"
	"quotes/pkg/app"
	"quotes/pkg/logger"
)

func main() {
	// TODO: logger init
	log := logger.LoggerInit()

	// TODO: load config:
	cfg, err := config.ConfigInit()
	if err != nil {
		log.Error("failed to load config", "err", err)
		return
	}

	err = app.StartApp(cfg)
	if err != nil {
		log.Error("failed to start", "err", err)
		return
	}

}
