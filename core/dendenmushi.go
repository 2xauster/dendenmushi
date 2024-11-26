package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jauster101/dendenmushi/commands"
	"github.com/jauster101/dendenmushi/core/logger"
	"github.com/jauster101/dendenmushi/handlers"
	"github.com/zekrotja/ken"
)

type DenDenMushi struct {
	session *discordgo.Session
	k     *ken.Ken
}

func must(err error) {
	if err != nil {
		logger.Err(fmt.Errorf("must failed: %v", err))
		os.Exit(1)
	}
}

func NewDenDenMushi() *DenDenMushi {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		must(fmt.Errorf("environmental variable DISCORD_TOKEN is empty"))
	}

	s, err := discordgo.New("Bot " + token)
	must(err)

	k, kErr := ken.New(s)
	must(kErr)

	return &DenDenMushi{
		session: s,
		k: k,
	}
}

func (ddm *DenDenMushi) LoadCommands() {
	logger.Info("Loading commands...")
	err := ddm.k.RegisterCommands(
		new(commands.PingCommand),
	)
	must(err)
	logger.Info("Commands loaded")
}

func (ddm *DenDenMushi) LoadHandlers() {
	logger.Info("Loading handlers...")
	ddm.session.AddHandler(handlers.ReadyHandler)
}
func (ddm *DenDenMushi) Start() {
	logger.Info("Starting app...")

	sErr := ddm.session.Open()
	must(sErr)

	logger.Info("Discord session opened. Press Ctrl + C to close.")

	defer ddm.session.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<- stop
	logger.Info("App closed")
}