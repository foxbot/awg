package awg

import (
	"encoding/json"
	"time"

	"github.com/foxbot/awg/wumpus"
	"github.com/go-redis/redis"
)

const (
	// KeyExchange is the key for the events exchange
	KeyExchange = "exchange:events"
)

// IWorker defines an interface for a worker
type IWorker interface {
	Discord() *wumpus.Discord
	Messages() <-chan wumpus.Message
}

// Worker will pull and parse data from Redis
type Worker struct {
	client   *redis.Client
	discord  *wumpus.Discord
	messages chan wumpus.Message
}

// NewWorker creates a new worker at the given redis address
func NewWorker(redisAddr, discordAddr string) (*Worker, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	discord := wumpus.NewDiscord(discordAddr)

	messages := make(chan wumpus.Message, 16)
	return &Worker{
		client:   client,
		discord:  discord,
		messages: messages,
	}, nil
}

// Close closes the worker
func (worker *Worker) Close() {
	close(worker.messages)

	worker.client.Close()
}

func (worker *Worker) Messages() <-chan wumpus.Message {
	return worker.messages
}

// Run runs the worker
func (worker *Worker) Run() <-chan error {
	errChan := make(chan error)
	go worker.run(errChan)
	return errChan
}
func (worker *Worker) run(errChan chan<- error) {
	for {
		result, err := worker.client.BLPop(time.Duration(0), KeyExchange).Result()
		if err != nil {
			errChan <- err
			return
		}

		b := []byte(result[1])
		ev := new(wumpus.Event)

		err = json.Unmarshal(b, &ev)
		if err != nil {
			errChan <- err
			return
		}

		switch ev.Type {
		case "MESSAGE_RECEIVE":
			err = worker.messageReceived(ev)
		}

		if err != nil {
			errChan <- err
			return
		}
	}
}

func (worker *Worker) messageReceived(ev *wumpus.Event) error {
	var m wumpus.Message
	err := json.Unmarshal(ev.Data, &m)
	if err != nil {
		return err
	}
	worker.messages <- m

	return nil
}
