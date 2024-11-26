package main

import (
	"fmt"

	"github.com/jauster101/dendenmushi/core"
	"github.com/jauster101/dendenmushi/core/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Err(fmt.Errorf("failed to load .env file: %v", err))
		return
	}

	ddm := core.NewDenDenMushi()
	ddm.LoadCommands()
	ddm.Start()
}