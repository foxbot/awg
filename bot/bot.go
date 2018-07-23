package bot

import (
	"github.com/foxbot/awg"
	"github.com/foxbot/awg/commands"
	"github.com/foxbot/awg/wumpus"
)

// Manager contains this bot's Commands
var Manager commands.Commands

// Bot contains behavior to manage commands
type Bot struct {
	Worker *awg.Worker
}

// Command will ask the bot to handle a command from a raw message
func (b *Bot) Command(msg wumpus.Message) error {
	// TODO: tokenizer/parsing
	ctx := &commands.Context{
		Message: msg,
		Worker:  b.Worker,
	}
	err := Manager.Invoke(ctx)
	return err
}
