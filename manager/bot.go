package manager

import "github.com/haashemi/tgo"

// Bot holds essential and critical fields about our bots.
//
// It is also totally safe to store the tokens here as we already
// have them in the telegram-bot-api's dir. so don't worry.
type Bot struct {
	ID       int64  `json:"id"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

// NewBot verifies the token/bot and returns a *Bot if everything goes fine.
func NewBot(token, telegramBotApiHost string) (*Bot, error) {
	botClient := tgo.NewAPI(token, telegramBotApiHost, nil)

	info, err := botClient.GetMe()
	if err != nil {
		return nil, err
	}

	return &Bot{
		ID:       info.Id,
		Token:    token,
		Username: info.Username,
	}, nil
}
