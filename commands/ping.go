package commands

import (
	"gobot/models"
)

var PingCommand = models.Command{
	Name:          "ping",
	Desc:          "Shows the bots ping.",
	Aliases:       []string{"p"},
	Args:          nil,
	Subcommands:   []string{""},
	Parentcommand: "none",
	Checks:        []func(*models.Context) error{},
	Callback:      pingCommand,
	Nsfw:          false,
	Endpoint:      "string",
}

func pingCommand(ctx *models.Context, args []string) {
	ctx.Send("pong!")
}
