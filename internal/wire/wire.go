//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Pugma/slack-traq-mcp/internal/config"
	"github.com/Pugma/slack-traq-mcp/internal/litellm"
	"github.com/Pugma/slack-traq-mcp/internal/mcp"
	"github.com/Pugma/slack-traq-mcp/internal/orchestrator"
	"github.com/Pugma/slack-traq-mcp/internal/slack"
	"github.com/Pugma/slack-traq-mcp/internal/traq"
	"github.com/google/wire"
)

var providers = wire.NewSet(
	config.NewConfig,
	traq.NewBot,
	slack.NewClient,
	litellm.NewClient,
	orchestrator.NewService,
	mcp.NewClient,
)

func InitializeApp() (*orchestrator.Service, error) {
	wire.Build(providers)
	return nil, nil
}
