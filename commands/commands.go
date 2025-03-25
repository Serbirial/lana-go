package commands

import (
	"gobot/models"
)

var AllCommands = map[string]models.Command{
	PingCommand.Name:      PingCommand,
	OwnerTestCommand.Name: OwnerTestCommand,
}
