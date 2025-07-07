package slack

import (
	"fmt"

	"github.com/Pugma/slack-traq-mcp/internal/config"
	"github.com/slack-go/slack"
)

type Client struct {
	api *slack.Client
}

type SlackChannel struct {
	ID   string
	Name string
}

func NewClient(cfg *config.Config) *Client {
	api := slack.New(cfg.SlackToken)
	return &Client{api}
}

// チャンネル ID とチャンネル名の一覧を取得
func (c *Client) GetChannels() ([]SlackChannel, error) {
	cursor := ""
	channels := []SlackChannel{}

	for {
		ch, nextCursor, err := c.api.GetConversations(&slack.GetConversationsParameters{
			Types: []string{"public_channel", "private_channel"},
		})
		if err != nil {
			return nil, fmt.Errorf("slack: failed to get channels: %w", err)
		}

		for _, channel := range ch {
			channels = append(channels, SlackChannel{
				ID:   channel.ID,
				Name: channel.Name,
			})
		}

		if cursor == "" {
			break
		}

		cursor = nextCursor
	}

	return channels, nil
}

// チャンネルのメッセージ履歴を取得
func (c *Client) GetChannelHistory(channelID string, count int) ([]slack.Message, error) {
	var allMessages []slack.Message
	cursor := ""

	for {
		history, err := c.api.GetConversationHistory(&slack.GetConversationHistoryParameters{
			ChannelID: channelID,
			Limit:     count,
			Cursor:    cursor,
		})
		if err != nil {
			return nil, fmt.Errorf("slack: failed to get channel history for %s: %w", channelID, err)
		}

		for _, message := range history.Messages {
			if message.Type == "message" {
				allMessages = append(allMessages, message)
			}
		}

		if history.ResponseMetadata.Cursor == "" {
			break
		}

		cursor = history.ResponseMetadata.Cursor
	}

	return allMessages, nil
}
