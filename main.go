package main

import (
	"log"

	"github.com/haashemi/BotManager/api"
	"github.com/haashemi/BotManager/bot"
	"github.com/haashemi/BotManager/config"
	"github.com/haashemi/BotManager/manager"
)

func main() {
	config, err := config.ParseConfig()
	if err != nil {
		log.Fatalln("Failed to parse config", err)
	}

	manager, err := manager.NewManager(config)
	if err != nil {
		log.Fatalln("Failed to initialize the bot manager", err)
	}

	runApi, err := api.NewAPI(config, manager)
	if err != nil {
		log.Fatalln("Failed to initialize the api", err)
	}

	runBot, err := bot.NewBot(config, manager)
	if err != nil {
		log.Fatalln("Failed to initialize the bot", err)
	}

	go func() { log.Fatalln(runApi()) }()
	go func() { log.Fatalln(runBot()) }()

	select {}
}
