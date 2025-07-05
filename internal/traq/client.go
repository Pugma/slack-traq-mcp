package traq

import (
	"github.com/Pugma/slack-traq-mcp/internal/config"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

type Bot struct {
	bot            *traqwsbot.Bot
	messageHandler func(channelID string, message string)
}

func NewBot(cfg *config.Config) (*Bot, error) {
	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: cfg.TraqBotAccessToken,
		Origin:      cfg.TraqWsURL,
	})
	if err != nil {
		return nil, err
	}

	return &Bot{bot, nil}, nil
}

func (b *Bot) Start() error {
	if err := b.bot.Start(); err != nil {
		return err
	}
	return nil
}
