package commands

import (
	"fmt"
	"log"

	"github.com/foxbot/awg/wumpus"
)

// Result is a contract for a command result
type Result interface {
	Act(ctx *Context) error
}

// ErrorResult is a result for a command that threw an error
type ErrorResult struct {
	Error error
}

// Error creates an ErrorResult
func Error(e error) ErrorResult {
	return ErrorResult{
		Error: e,
	}
}

// Act logs the error and writes a message to the channel
func (r ErrorResult) Act(ctx *Context) error {
	log.Println("A command passed up an error", r.Error)

	text := fmt.Sprintf("**Error:** %s", r.Error.Error())
	args := wumpus.CreateMessageArgs{
		Content: text,
	}
	_, err := ctx.Worker.Discord().CreateMessage(ctx.Message.ChannelID, args)
	return err
}

// TextResult is a result for a command that writes text
type TextResult struct {
	Message string
}

// Text creates a TextResult
func Text(msg string) TextResult {
	return TextResult{
		Message: msg,
	}
}

func Textf(format string, args ...interface{}) TextResult {
	msg := fmt.Sprintf(format, args...)
	return TextResult{
		Message: msg,
	}
}

// Act writes a message to the channel
func (r TextResult) Act(ctx *Context) error {
	args := wumpus.CreateMessageArgs{
		Content: r.Message,
	}
	_, err := ctx.Worker.Discord().CreateMessage(ctx.Message.ChannelID, args)
	return err
}
