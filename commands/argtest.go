package commands

import (
	"gobot/models"
	"gobot/utils/checks"
	"strconv"
	"strings"
)

var ArgTestCommand = models.Command{
	Name:          "argtest",
	Desc:          "Tests the arg system.",
	Aliases:       []string{"at"},
	Args:          map[string]string{"arg1": "Testing.", "arg2": "Testing 2"},
	Subcommands:   []string{""},
	Parentcommand: "none",
	Checks:        []func(*models.Context) error{checks.IsOwner, checks.NeedsArgs},
	Callback:      argTestCommand,
	Nsfw:          false,
	Endpoint:      "string",
}

func argTestCommand(ctx *models.Context, argsRaw string) {
	var toSend = ""
	var args []string = strings.Fields(argsRaw)

	for i := 0; i < len(args); i++ {
		toSend += "Arg **" + strconv.Itoa(i) + "**: " + args[i] + "\n"
	}
	ctx.Send(toSend)
}
