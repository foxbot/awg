package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/foxbot/awg"
	"github.com/foxbot/awg/bot"
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

	blazeAddr := os.Getenv("BLAZE_ADDR")
	if blazeAddr == "" {
		panic("missing BLAZE_ADDR")
	}

	worker, err := awg.NewWorker(redisAddr, blazeAddr)
	if err != nil {
		panic(err)
	}
	defer worker.Close()

	bot := bot.Bot{
		Worker: worker,
	}

	errChan := worker.Run()
	msgChan := worker.Messages()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

loop:
	for {
		select {
		case msg := <-msgChan:
			err := bot.Command(msg)
			if err != nil {
				// TODO: should bot error trash the worker?
				panic(err)
			}
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
