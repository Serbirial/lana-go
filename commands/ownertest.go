package commands

import (
	"gobot/models"
	"gobot/utils/checks"
)

var OwnerTestCommand = models.Command{
	Name:          "ownertest",
	Desc:          "Shouldnt be runnable by normal users.",
	Aliases:       []string{"ot"},
	Args:          nil,
	Subcommands:   []string{""},
	Parentcommand: "none",
	Checks:        []func(*models.Context) error{checks.IsOwner},
	Callback:      test,
	Nsfw:          false,
	Endpoint:      "string",
}

func test(ctx *models.Context, argsRaw string) {
	ctx.Send("You are an owner.")
}
