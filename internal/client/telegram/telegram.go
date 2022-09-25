package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var token = ""

type Client interface {
	SendMessage(userID int64, messageText string) error
	GetUpdates() (map[int64][]string, error)
}

type client struct {
	bot    *tgbotapi.BotAPI
	offset int64
}

var _ Client = &client{}

// mb use https://github.com/gotd/td ?
func NewClient() Client {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil
	}
	return &client{
		bot:    bot,
		offset: 0,
	}
}

func (c *client) SendMessage(userID int64, txt string) error {
	return nil
}

func (c *client) GetUpdates() (map[int64][]string, error) {
	u := tgbotapi.NewUpdate(c.offset)
	updates := c.bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
	return nil, nil
}
