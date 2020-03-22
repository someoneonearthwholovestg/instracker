package box

import (
	"Instracker/internal/config"
	"Instracker/internal/database"
	"Instracker/internal/instagram"
	"Instracker/internal/telegrambotapi"
)

type Box struct {
	Config         *config.Config
	Database       *database.Database
	Instagram      *instagram.Instagram
	TelegramBotAPI *telegrambotapi.TelegramBotAPI
}

// NewBox wire provider
func NewBox(
	conf *config.Config,
	db *database.Database,
	instagram *instagram.Instagram,
	tgbotapi *telegrambotapi.TelegramBotAPI,
) *Box {
	return &Box{
		Config:         conf,
		Database:       db,
		Instagram:      instagram,
		TelegramBotAPI: tgbotapi,
	}
}
