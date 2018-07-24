package commands

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
	name := ctx.Message.Content[offset:] // TODO: tokenize/parse
	println("name=", name)

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
