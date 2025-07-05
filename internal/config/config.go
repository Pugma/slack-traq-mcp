package config

import "github.com/caarlos0/env/v11"

type Config struct {
	OpenAIURL          string `env:"OPENAI_URL"`
	OpenAIToken        string `env:"OPENAI_TOKEN"`
	SlackToken         string `env:"SLACK_TOKEN"`
	TraqBotAccessToken string `env:"TRAQ_BOT_ACCESS_TOKEN"`
	TraqWsURL          string `env:"TRAQ_WS_URL"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
