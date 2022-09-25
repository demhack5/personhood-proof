package telegram

import (
	"log"
	"personhood-proof/internal/models/telegram"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var token = ""

type Client interface {
	SendMessage(msg *telegram.Message) error
	GetUpdates() (map[int64][]*telegram.Message, error)
}

type client struct {
	bot    *tgbotapi.BotAPI
	offset int
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

func (c *client) SendMessage(msg *telegram.Message) error {
	m := tgbotapi.NewMessage(msg.ChatID, msg.Text)
	c.bot.Send(m)
	return nil
}

func (c *client) GetUpdates() (map[int64][]*telegram.Message, error) {
	ret := make(map[int64][]*telegram.Message)
	u := tgbotapi.NewUpdate(c.offset)
	updates, err := c.bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			a, ok := ret[int64(update.Message.From.ID)]
			if !ok {
				ret[int64(update.Message.From.ID)] = make([]*telegram.Message, 0)
				ret[int64(update.Message.From.ID)] = append(ret[int64(update.Message.From.ID)], &telegram.Message{
					Text:   update.Message.Text,
					ChatID: update.Message.Chat.ID,
				})
			} else {
				ret[int64(update.Message.From.ID)] = append(a, &telegram.Message{
					Text:   update.Message.Text,
					ChatID: update.Message.Chat.ID,
				})
			}
		}
	}
	return ret, nil
}
