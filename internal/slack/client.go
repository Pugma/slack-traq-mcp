package slack

import (
	"github.com/Pugma/slack-traq-mcp/internal/config"
	"github.com/slack-go/slack"
)

type Client struct {
	api *slack.Client
}

func NewClient(cfg *config.Config) *Client {
	api := slack.New(cfg.SlackToken)
	return &Client{api}
}
