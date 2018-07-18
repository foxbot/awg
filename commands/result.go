package commands

import (
	"log"
)

// Result is a contract for a command result
type Result interface {
	Act(ctx *Context) error
}

// ErrorResult is a result for a command that threw an error
type ErrorResult struct {
	Error error
}

// Act logs the error and writes a message to the channel
func (r *ErrorResult) Act(ctx *Context) error {
	log.Println("A command passed up an error", r.Error)
	// TODO handle error
	return nil
}

// TextResult is a result for a command that writes text
type TextResult struct {
	Message string
}

// Act writes a message to the channel
func (r *TextResult) Act(ctx *Context) error {
	// TODO send message
	return nil
}
