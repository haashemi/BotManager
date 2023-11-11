package bot

import (
	"fmt"
	"strings"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
)

const onRemHelpText = `Send the token with the command.

Example:
<code>/rem 324230498:aabbccddeeff0011223344</code>
`

const onRemText = `Successfully removed @%s`

// onRem handles /rem command.
func (b *Bot) onRem(ctx *message.Context) {
	msg := strings.SplitN(ctx.String(), " ", 2)

	if len(msg) == 1 {
		ctx.Send(&tgo.SendMessage{Text: onRemHelpText})
		return
	}

	ctx.Delete()
	ctx.Bot.SendChatAction(&tgo.SendChatAction{ChatId: tgo.ID(ctx.From.Id), Action: "typing"})

	bot, err := b.manager.RemoveBot(msg[1])
	if err != nil {
		handleError(ctx, err, "Failed to remove the bot")
		return
	}

	ctx.Send(&tgo.SendMessage{Text: fmt.Sprintf(onRemText, bot.Username)})
}
