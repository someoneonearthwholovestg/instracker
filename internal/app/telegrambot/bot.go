package telegrambot

import (
	"Instracker/internal/app/telegrambot/db"
	"Instracker/internal/box"
	"Instracker/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"os"
)

type InstaBot struct {
	Box     *box.Box
	Answers map[string]string
}

const (
	stateZERO            = iota // 0
	stateLISTUNFOLLOWERS        // 1
	stateSUBSCRIBE              // 2
	stateUNSUBSCRIBE            // 3
)

func NewInstaBot(b *box.Box) (*InstaBot, error) {
	answers, err := pkg.GetMapFromJSON(b.Config.Dialogue)
	if err != nil {
		return nil, err
	}

	log.Info(b.Config.Vault)
	if _, err := os.Stat(b.Config.Vault); os.IsNotExist(err) {
		return nil, err
	}

	return &InstaBot{
		Box:     b,
		Answers: answers,
	}, nil
}

// Run ...
func (i *InstaBot) Run() {
	log.Info("telegram bot is starting")

	updates, err := i.Box.TelegramBotAPI.GetUpdatesChanel()
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go i.Manage(update)

		log.WithFields(log.Fields{
			"user_id":    update.Message.Chat.ID,
			"username":   update.Message.Chat.UserName,
			"first_name": update.Message.Chat.FirstName,
			"last_name":  update.Message.Chat.LastName,
			"message":    update.Message.Text,
		}).Info("MESSAGE")
	}
}

// Manage plays router role to set an appropriate state for user
func (i *InstaBot) Manage(update tgbotapi.Update) {
	var msg = update.Message.Text
	switch msg {
	case "/start":
		i.CommandsHandler(update, msg, stateZERO)
	case "/help":
		i.CommandsHandler(update, msg, stateZERO)
	case "/listunfollowers":
		i.CommandsHandler(update, msg, stateLISTUNFOLLOWERS)
	case "/subscribe":
		i.CommandsHandler(update, msg, stateSUBSCRIBE)
	case "/unsubscribe":
		i.CommandsHandler(update, msg, stateUNSUBSCRIBE)
	default:
		i.Handler(update, msg)
	}
}

// CommandsHandler processes format commands like "/{command}"
func (i *InstaBot) CommandsHandler(
	update tgbotapi.Update,
	msg string,
	state int,
) {
	i.Box.TelegramBotAPI.Send(update.Message.Chat.ID, i.Answers[msg])

	db.CreateOrUpdateUser(i.Box.Database.Conn, &db.User{
		ID:        update.Message.Chat.ID,
		Username:  update.Message.Chat.UserName,
		FirstName: update.Message.Chat.FirstName,
		LastName:  update.Message.Chat.LastName,
		State:     state,
	})
}

// Handler checks user state and then does some stuff for this state
func (i *InstaBot) Handler(
	update tgbotapi.Update,
	msg string,
) {
	i.Box.TelegramBotAPI.Send(update.Message.Chat.ID, msg)
}
