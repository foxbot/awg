package commands

import (
	"github.com/foxbot/awg"
	"github.com/foxbot/awg/wumpus"
)

// Context contains the context for a command
type Context struct {
	Message wumpus.Message
	Worker  *awg.Worker
}
