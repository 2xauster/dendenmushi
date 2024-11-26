package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jauster101/dendenmushi/core/logger"
	"github.com/zekrotja/ken"
)

type PingCommand struct{}

var _ ken.SlashCommand = (*PingCommand)(nil)

func (c *PingCommand) Name() string {
	return "ping"
}

func (c *PingCommand) Description() string {
	return "Shows the bot's latency"
}

func (c *PingCommand) Version() string {
	return "0.2"
}

func (c *PingCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *PingCommand) Run(ctx ken.Context) (err error) {
	session := ctx.GetSession()

	start := time.Now()
	_, err = session.User("@me")
	if err != nil {
		logger.Err(fmt.Errorf("command %s failed: %v", c.Name(), err))
		return ctx.RespondError("Something went wrong", "Internal error")
	}
	apiLatency := time.Since(start)

	heartbeatLatency := session.HeartbeatLatency()

	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				"💗 Heartbeat latency: **`%v`**\n📡 API latency: **`%v`**",
				heartbeatLatency,
				apiLatency,
			),
		},
	})
	return
}
