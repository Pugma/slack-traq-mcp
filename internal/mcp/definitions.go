package mcp

import (
	"context"

	"github.com/Pugma/slack-traq-mcp/internal/slack"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/server"
)

type Client struct {
	mcpClient *client.Client
}

func NewClient(slackClient *slack.Client) (*Client, error) {
	mcpClient, err := client.NewInProcessClient((server.NewMCPServer("", "")))
	if err != nil {
		return nil, err
	}

	return &Client{mcpClient}, nil
}

func (c *Client) Start() error {
	if err := c.mcpClient.Start(context.Background()); err != nil {
		return err
	}
	return nil
}
