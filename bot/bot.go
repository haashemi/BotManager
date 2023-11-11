package bot

import (
	"github.com/haashemi/BotManagerBot/config"
	"github.com/haashemi/BotManagerBot/manager"
	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/filters"
	"github.com/haashemi/tgo/routers/message"
)

type RunFunc func() error

type Bot struct {
	telegramBotApiHost string

	manager *manager.Manager
}

// NewBot initializes a Bot instance and returns the StartPolling method
func NewBot(config *config.Config, manager *manager.Manager) (RunFunc, error) {
	bot := tgo.NewBot(config.Bot.Token, tgo.Options{Host: config.TelegramBotAPI.Host, DefaultParseMode: tgo.ParseModeHTML})

	info, err := bot.GetMe()
	if err != nil {
		return nil, err
	}

	instance := &Bot{telegramBotApiHost: config.TelegramBotAPI.Host, manager: manager}
	whitelistFilter := filters.Whitelist(config.Bot.Admins...)

	mr := message.NewRouter()
	mr.Handle(filters.And(filters.Command("add", info.Username), whitelistFilter), instance.onAdd)
	mr.Handle(filters.And(filters.Command("rem", info.Username), whitelistFilter), instance.onRem)
	mr.Handle(filters.And(filters.Command("bots", info.Username), whitelistFilter), instance.onBots)
	bot.AddRouter(mr)

	bot.SetMyCommands(&tgo.SetMyCommands{
		Commands: []*tgo.BotCommand{
			{Command: "add", Description: "Whitelist a new bot"},
			{Command: "rem", Description: "Remove a bot from the whitelist"},
			{Command: "bots", Description: "List of whitelisted bots"},
		},
	})

	return func() error { return bot.StartPolling(30) }, nil
}
