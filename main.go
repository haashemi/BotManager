package main

import (
	"log"

	"github.com/haashemi/BotManagerBot/api"
	"github.com/haashemi/BotManagerBot/bot"
	"github.com/haashemi/BotManagerBot/config"
	"github.com/haashemi/BotManagerBot/manager"
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

	go log.Fatalln(runApi())
	go log.Fatalln(runBot())

	select {}
}
