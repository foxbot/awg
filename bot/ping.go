package bot

import (
	"github.com/foxbot/awg/commands"
)

func init() {
	pingCommand := commands.Command{
		Aliases: []string{"ping", "pong", "hello"},
		Func:    ping,
	}
	Manager.Add(pingCommand)
}

func ping(ctx *commands.Context) commands.Result {
	return commands.Textf("pong! from worker %s", ctx.Worker.ID)
}
