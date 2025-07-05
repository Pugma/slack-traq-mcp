package litellm

import (
	"github.com/Pugma/slack-traq-mcp/internal/config"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Client struct {
	client *openai.Client
}

func NewClient(cfg *config.Config) *Client {
	client := openai.NewClient(
		option.WithBaseURL(cfg.OpenAIURL),
		option.WithAPIKey(cfg.OpenAIToken),
	)
	return &Client{
		client: &client,
	}
}
