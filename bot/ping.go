package bot

import (
	"github.com/dabbotorg/worker/commands"
)

func init() {
	pingCommand := commands.Command{
		Aliases: []string{"ping", "pong", "hello"},
		Func:    ping,
	}
	Manager.Add(pingCommand)
}

func ping(ctx *commands.Context) commands.Result {
	return commands.Text("pong!")
}
