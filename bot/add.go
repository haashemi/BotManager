package bot

import (
	"fmt"
	"strings"

	"github.com/haashemi/BotManagerBot/manager"
	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
)

const onAddHelpText = `Send the token with the command.

Example:
<code>/add 324230498:aabbccddeeff0011223344</code>
`

const onAddText = `Successfully whitelisted @%s`

// onAdd handles /add command
func (b *Bot) onAdd(ctx *message.Context) {
	msg := strings.SplitN(ctx.String(), " ", 2)

	if len(msg) == 1 {
		ctx.Send(&tgo.SendMessage{Text: onAddHelpText})
		return
	}

	ctx.Delete()
	ctx.Bot.SendChatAction(&tgo.SendChatAction{ChatId: tgo.ID(ctx.From.Id), Action: "typing"})

	bot, err := manager.NewBot(msg[1], b.telegramBotApiHost)
	if err != nil {
		handleError(ctx, err, "Failed to verify the bot")
		return
	}

	if err := b.manager.AddBot(bot); err != nil {
		handleError(ctx, err, "Failed to add the bot")
		return
	}

	ctx.Send(&tgo.SendMessage{Text: fmt.Sprintf(onAddText, bot.Username)})
}
