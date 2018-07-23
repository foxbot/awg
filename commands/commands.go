package commands

import "strings"

const (
	// Prefix is the bot's prefix (TODO: dynamic)
	Prefix = "~>>"
)

// Executor defines the delegate for a command
type Executor func(ctx *Context) Result

// Command is a command that may be invoked
type Command struct {
	Aliases []string
	Func    Executor
}

// Commands is a structure to manage commands
type Commands struct {
	commands []Command
}

// Add registers a command with the Commands manager
func (c *Commands) Add(command Command) {
	c.commands = append(c.commands, command)
}

// Invoke runs a command
func (c *Commands) Invoke(ctx *Context) error {
	if !strings.HasPrefix(ctx.Message.Content, "~>>") {
		return nil
	}
	name := ctx.Message.Content[len(Prefix)-1:] // TODO: tokenize/parse

	var command *Command
	for _, cmd := range c.commands {
		for _, alias := range cmd.Aliases {
			if name == alias {
				command = &cmd
				break
			}
		}
	}

	if command == nil {
		// TODO: command not found?
		return nil
	}

	result := command.Func(ctx)
	err := result.Act(ctx)
	return err
}
