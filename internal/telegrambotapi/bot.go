package telegrambotapi

import (
	"Instracker/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewTelegramBotAPI create an instance of Bot
func NewTelegramBotAPI(cfg *config.Config) (*TelegramBotAPI, error) {
	api, err := tgbotapi.NewBotAPI(cfg.TelegramBot.Token)
	if err != nil {
		return nil, err
	}

	return &TelegramBotAPI{api}, nil
}

// Bot is a custom wrapper for Telegram Bot
type TelegramBotAPI struct {
	API *tgbotapi.BotAPI
}

// GetUpdatesChanel returns UpdatesChannel for receiving new updates
func (i *TelegramBotAPI) GetUpdatesChanel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := i.API.GetUpdatesChan(u)
	return updates, err
}

// Send sends message to user
func (i *TelegramBotAPI) Send(userID int64, text string) (err error) {
	msg := tgbotapi.NewMessage(userID, text)
	_, err = i.API.Send(msg)
	return
}

// SendDocument sends a document to user
func (i *TelegramBotAPI) SendDocument(userID int64, file interface{}) (err error) {
	msg := tgbotapi.NewDocumentUpload(userID, file)
	_, err = i.API.Send(msg)
	return
}
