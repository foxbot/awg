package commands

import (
	"github.com/dabbotorg/worker"
	"github.com/dabbotorg/worker/wumpus"
)

// Context contains the context for a command
type Context struct {
	Message wumpus.Message
	Worker  worker.Worker
}
