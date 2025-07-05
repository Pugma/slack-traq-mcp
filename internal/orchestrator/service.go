package orchestrator

import (
	"fmt"
	"log/slog"

	"github.com/Pugma/slack-traq-mcp/internal/litellm"
	"github.com/Pugma/slack-traq-mcp/internal/mcp"
	"github.com/Pugma/slack-traq-mcp/internal/traq"
)

type Service struct {
	llmClient *litellm.Client
	mcpClient *mcp.Client
	traqBot   *traq.Bot
}

func NewService(llmClient *litellm.Client, mcpClient *mcp.Client, traqBot *traq.Bot) *Service {
	return &Service{
		llmClient,
		mcpClient,
		traqBot,
	}
}

func (s *Service) Start() error {
	go func() {
		if err := s.mcpClient.Start(); err != nil {
			slog.Error(fmt.Sprintf("failed to start mcp client: %v", err))
		}
	}()

	if err := s.traqBot.Start(); err != nil {
		return err
	}

	return nil
}
