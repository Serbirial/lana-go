package commands

import (
	"gobot/models"
	"gobot/utils/checks"
	"strconv"
)

var ArgTestCommand = models.Command{
	Name:          "argtest",
	Desc:          "Tests the arg system.",
	Aliases:       []string{"at"},
	Subcommands:   []string{""},
	Parentcommand: "none",
	Checks:        []func(*models.Context) error{checks.IsOwner},
	Callback:      argTestCommand,
	Nsfw:          false,
	Endpoint:      "string",
}

func argTestCommand(ctx *models.Context, args []string) {
	var toSend = ""
	for i := 0; i < len(args); i++ {
		toSend += "Arg **" + strconv.Itoa(i) + "**: " + args[i] + "\n"
	}
	ctx.Send(toSend)
}
