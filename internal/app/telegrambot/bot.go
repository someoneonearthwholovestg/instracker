package telegrambot

import (
	"Instracker/internal/box"
	"Instracker/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"

	log "github.com/sirupsen/logrus"
)

type InstaBot struct {
	Box     *box.Box
	Answers map[string]string
}

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

		go i.Manager(update)

		log.WithFields(log.Fields{
			"user_id":    update.Message.Chat.ID,
			"username":   update.Message.Chat.UserName,
			"first_name": update.Message.Chat.FirstName,
			"last_name":  update.Message.Chat.LastName,
			"message":    update.Message.Text,
		}).Info("new request")
	}
}

// Manager is a router for message handlers
func (i *InstaBot) Manager(update tgbotapi.Update) {
	if err := i.Box.TelegramBotAPI.Send(update.Message.Chat.ID, "hi"); err != nil {
		log.WithError(err)
	}
	//switch update.Message.Text {
	//case "/start":
	//	i.commonHandler(update, update.Message.Text, stateZERO)
	//case "/help":
	//	i.commonHandler(update, update.Message.Text, stateZERO)
	//case "/listunfollowers":
	//	i.commonHandler(update, update.Message.Text, stateLISTUNFOLLOWERS)
	//case "/subscribe":
	//	i.commonHandler(update, update.Message.Text, stateSUBSCRIBE)
	//case "/unsubscribe":
	//	i.commonHandler(update, update.Message.Text, stateUNSUBSCRIBE)
	//case "/cancel":
	//	i.cancelHandler(update)
	//default:
	//	i.statesHandler(update)
	//}
}
