package bot

import (
	"fmt"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
)

const errorMessage = `🚫| Error Occurred.

⚠️| %s

ℹ️| %s
`

func handleError(ctx *message.Context, err error, msg string) {
	ctx.Send(&tgo.SendMessage{Text: fmt.Sprintf(errorMessage, msg, err.Error())})
}
