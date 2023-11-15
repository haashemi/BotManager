package bot

import (
	"fmt"
	"strings"

	"github.com/haashemi/BotManager/manager"
	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
	"github.com/samber/lo"
)

const onBotsText = `👑| Whitelisted bots are:

%s
`

const botInfoText = `— @%s [<code>%d</code>]
<span class="tg-spoiler"><code>%s</code></span>`

// onBots handles /bots command
func (b *Bot) onBots(ctx *message.Context) {
	bots := lo.Map(b.manager.Bots(), func(bot *manager.Bot, index int) string {
		return fmt.Sprintf(botInfoText, bot.Username, bot.ID, bot.Token)
	})

	ctx.Send(&tgo.SendMessage{Text: fmt.Sprintf(onBotsText, strings.Join(bots, "\n\n"))})
}
