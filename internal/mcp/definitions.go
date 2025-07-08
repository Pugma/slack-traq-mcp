package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Pugma/slack-traq-mcp/internal/slack"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Client struct {
	mcpClient *client.Client
}

func setupMCPServer(slackClient *slack.Client) *server.MCPServer {
	s := server.NewMCPServer(
		"", "",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)

	s.AddTool(
		mcp.NewTool(
			"get_slack_channels",
			mcp.WithDescription("Slackのチャンネル一覧を取得する"),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			channels, err := slackClient.GetChannels()
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := json.Marshal(channels)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			return mcp.NewToolResultText(string(resp)), nil
		},
	)

	s.AddTool(
		mcp.NewTool(
			"get_slack_channel_history",
			mcp.WithDescription("指定したSlackチャンネルのメッセージ履歴を取得する"),
			mcp.WithString("channel_id",
				mcp.Required(),
				mcp.Description("取得対象のSlackチャンネルのチャンネルID"),
			),
			mcp.WithNumber("count",
				mcp.Required(),
				mcp.Description("取得するメッセージ数"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			ch_id, err := request.RequireString("channel_id")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			cnt, err := request.RequireInt("count")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			messages, err := slackClient.GetChannelHistory(ch_id, cnt)
			if err != nil {
				return nil, fmt.Errorf("failed to get channel history: %w", err)
			}
			resp, err := json.Marshal(messages)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal messages: %w", err)
			}

			return mcp.NewToolResultText(string(resp)), nil
		},
	)

	return s
}

func NewClient(slackClient *slack.Client) (*Client, error) {
	s := setupMCPServer(slackClient)

	mcpClient, err := client.NewInProcessClient(s)
	if err != nil {
		return nil, fmt.Errorf("mcp: failed to create client: %w", err)
	}

	return &Client{mcpClient}, nil
}

func (c *Client) Start() error {
	_, err := c.mcpClient.Initialize(context.Background(), mcp.InitializeRequest{})
	if err != nil {
		return err
	}

	return nil
}
