package main

import (
	"github.com/makves-test-task/api"
	"github.com/makves-test-task/config"
	"github.com/makves-test-task/helper"
	"github.com/makves-test-task/pkg/logger"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogLevel)

	fm, err := helper.New(log, &cfg)
	if err != nil {
		log.Fatal("failed to get file manager err: " + err.Error())
		return
	}
	server := api.New(api.Option{
		Conf:        cfg,
		Logger:      log,
		FileManager: fm,
	})

	if err := server.Run(":" + cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", err)
		panic(err)
	}
}
