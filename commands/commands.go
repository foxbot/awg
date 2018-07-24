package commands

import "strings"

const (
	// Prefix is the bot's prefix (TODO: dynamic)
	Prefix = "~>>"
)

// Executor defines the delegate for a command
type Executor func(ctx *Context) Result

// PrefixFunc defines the delegate to check a prefix
type PrefixFunc func(ctx *Context) (int, bool)

// Command is a command that may be invoked
type Command struct {
	Aliases []string
	Func    Executor
}

// Commands is a structure to manage commands
type Commands struct {
	commands   []Command
	prefixFunc PrefixFunc
}

// NewCommands returns a Commands manager
func NewCommands(prefixFunc PrefixFunc) *Commands {
	return &Commands{
		prefixFunc: prefixFunc,
	}
}

// Add registers a command with the Commands manager
func (c *Commands) Add(command Command) {
	c.commands = append(c.commands, command)
}

// Invoke runs a command
func (c *Commands) Invoke(ctx *Context) error {
	offset, ok := c.prefixFunc(ctx)
	if !ok {
		return nil
	}
	content := ctx.Message.Content[offset:] // TODO: tokenize/parse
	parts := strings.SplitN(content, " ", 2)
	name := parts[0]
	println("name=", name)
	ctx.Name = name

	if len(parts) > 1 {
		ctx.Argument = parts[1]
	}

	var command *Command

	for _, cmd := range c.commands {
		for _, alias := range cmd.Aliases {
			if name == alias {
				command = &cmd
				goto found
			}
		}
	}

	if command == nil {
		// TODO: command not found?
		return nil
	}
found:

	result := command.Func(ctx)
	err := result.Act(ctx)
	return err
}
