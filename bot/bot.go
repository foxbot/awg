// Package bot provides an example bot with example commands using this worker
package bot

import (
	"log"
	"strings"

	"github.com/foxbot/awg"
	"github.com/foxbot/awg/commands"
	"github.com/foxbot/awg/wumpus"
)

// Manager contains this bot's Commands
var Manager = commands.NewCommands(checkPrefix)

const prefix = "~>>"

func checkPrefix(ctx *commands.Context) (int, bool) {
	if strings.HasPrefix(ctx.Message.Content, prefix) {
		return len(prefix), true
	}
	return 0, false
}

// Bot contains behavior to manage commands
type Bot struct {
	Worker *awg.Worker
}

// Command will ask the bot to handle a command from a raw message
func (b *Bot) Command(msg wumpus.Message) error {
	log.Println("trying to run content=", msg.Content)
	// TODO: tokenizer/parsing
	ctx := &commands.Context{
		Message: msg,
		Worker:  b.Worker,
	}
	err := Manager.Invoke(ctx)
	return err
}
