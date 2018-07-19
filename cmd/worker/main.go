package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/dabbotorg/worker"
	"github.com/dabbotorg/worker/bot"
	"github.com/joho/godotenv"
)

var (
	id string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// TODO: different worker ID if using docker?
	id = strconv.Itoa(os.Getpid())
}

func main() {
	log.Printf("worker %s up\n", id)

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		panic("missing REDIS_ADDR")
	}

	worker, err := worker.NewWorker(redisAddr)
	if err != nil {
		panic(err)
	}
	defer worker.Close()

	errChan := worker.Run()
	msgChan := worker.Messages

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

loop:
	for {
		select {
		case <-msgChan:
			break
		case err = <-errChan:
			if err != nil {
				panic(err)
			}
		case <-sigChan:
			break loop
		}
	}

	log.Printf("worker %s down\n", id)
}
