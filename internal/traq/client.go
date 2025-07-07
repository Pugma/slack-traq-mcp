package traq

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Pugma/slack-traq-mcp/internal/config"
	"github.com/traPtitech/go-traq"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
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
		return nil, fmt.Errorf("traq: failed to initialize bot: %w", err)
	}

	return &Bot{bot, nil}, nil
}

func (b *Bot) SetMessageHandler(handler func(channelID string, message string)) {
	b.messageHandler = handler
}

func (b *Bot) Start() error {
	b.bot.OnMessageCreated(func(p *payload.MessageCreated) {
		if p == nil {
			slog.Warn("traq: received nil payload in OnMessageCreated")
			return
		}

		b.messageHandler(p.Message.ChannelID, p.Message.PlainText)
	})

	if err := b.bot.Start(); err != nil {
		return fmt.Errorf("traq: failed to start bot: %w", err)
	}
	return nil
}

func (b *Bot) SendMessage(channelID string, content string) error {
	_, _, err := b.bot.API().MessageApi.
		PostMessage(context.Background(), channelID).
		PostMessageRequest(traq.PostMessageRequest{
			Content: content,
		}).
		Execute()
	if err != nil {
		return fmt.Errorf("traq: failed to send message to channel %s: %v", channelID, err)
	}
	return nil
}
