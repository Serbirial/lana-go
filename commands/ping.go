package commands

import (
	"gobot/models"
)

var PingCommand = models.Command{
	Name:          "ping",
	Desc:          "Shows the bots ping.",
	Aliases:       []string{"p"},
	Subcommands:   []string{""},
	Parentcommand: "none",
	Checks:        []func(*models.Context) error{},
	Callback:      ping,
	Nsfw:          false,
	Endpoint:      "string",
}

func ping(ctx *models.Context) {
	ctx.Send("pong!")
}
