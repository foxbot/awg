package main

import (
	"encoding/json"
	"time"

	"github.com/dabbotorg/worker/wumpus"
	"github.com/go-redis/redis"
)

const (
	// KeyExchange is the key for the events exchange
	KeyExchange = "exchange:events"
)

// Worker will pull and parse data from Redis
type Worker struct {
	client *redis.Client
}

// NewWorker creates a new worker at the given redis address
func NewWorker(redisAddr string) (*Worker, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Worker{
		client: client,
	}, nil
}

// Run runs the worker
func (worker *Worker) Run() chan error {
	errChan := make(chan error)
	go worker.run(errChan)
	return errChan
}
func (worker *Worker) run(errChan chan error) {
	for {
		result, err := worker.client.BLPop(time.Duration(0), KeyExchange).Result()
		if err != nil {
			errChan <- err
			return
		}

		b := []byte(result[1])
		var ev wumpus.Event

		err = json.Unmarshal(b, &ev)
		if err != nil {
			errChan <- err
			return
		}
	}
}
