package bot

import (
	"github.com/dabbotorg/worker"
	"github.com/dabbotorg/worker/commands"
	"github.com/dabbotorg/worker/wumpus"
)

// Manager contains this bot's Commands
var Manager commands.Commands

// Bot contains behavior to manage commands
type Bot struct {
	Worker *worker.Worker
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
