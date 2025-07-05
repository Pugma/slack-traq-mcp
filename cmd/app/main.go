package main

import (
	"log"

	"github.com/Pugma/slack-traq-mcp/internal/wire"
)

func main() {
	// アプリケーションを初期化
	app, err := wire.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// アプリケーションのリッスンを開始
	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
