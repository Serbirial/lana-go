package commands

import (
	"gobot/models"
)

var AllCommands = map[string]models.Command{
	HelpCommand.Name:       HelpCommand,
	PingCommand.Name:       PingCommand,
	ArgTestCommand.Name:    ArgTestCommand,
	OwnerTestCommand.Name:  OwnerTestCommand,
	OwnersListCommand.Name: OwnersListCommand,
}
