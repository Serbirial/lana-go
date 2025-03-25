package commands

import (
	"gobot/models"
	"gobot/utils/checks"
)

var OwnerTestCommand = models.Command{
	Name:          "test",
	Desc:          "Shouldnt be runnable by normal users.",
	Aliases:       []string{"t"},
	Subcommands:   []string{""},
	Parentcommand: "none",
	Checks:        []func(*models.Context) error{checks.IsOwner},
	Callback:      ping,
	Nsfw:          false,
	Endpoint:      "string",
}

func test(ctx *models.Context) {
	ctx.Send("pong!")
}
