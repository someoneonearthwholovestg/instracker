//+build wireinject

package box

import (
	"Instracker/internal/config"
	"Instracker/internal/database"
	"Instracker/internal/instagram"
	"Instracker/internal/telegrambotapi"
	"github.com/google/wire"
)

// InitializeBox wire injection
func InitializeBox(configPath string) (*Box, error) {
	wire.Build(
		NewBox,
		config.NewConfig,
		database.NewDatabase,
		instagram.NewInstagram,
		telegrambotapi.NewTelegramBotAPI,
	)

	return nil, nil
}
