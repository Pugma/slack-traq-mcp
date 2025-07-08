package litellm

import (
	"context"
	"fmt"

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

// GenerateResponse takes a context and a prompt, and returns a response from the LLM.
func (c *Client) GenerateResponse(ctx context.Context, prompt string) (string, error) {
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT4o,
	}

	completion, err := c.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", fmt.Errorf("litellm: failed to generate response: %w", err)
	}

	if len(completion.Choices) == 0 {
		return "", fmt.Errorf("litellm: no response choices returned")
	}

	return completion.Choices[0].Message.Content, nil
}
