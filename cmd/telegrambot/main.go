package main

import (
	"Instracker/internal/app/telegrambot"
	"Instracker/internal/box"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)

	configPath, found := os.LookupEnv("CFG_PATH")
	if !found {
		print(errors.New("no config path variable"))
	}

	b, err := box.InitializeBox(configPath)
	if err != nil {
		panic(err)
	}

	tgbot, err := telegrambot.NewInstaBot(b)
	if err != nil {
		panic(err)
	}

	tgbot.Run()
}
