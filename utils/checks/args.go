package checks

import (
	"errors"
	"fmt"
	"gobot/models"
	"strconv"
	"strings"
)

func NeedsArgs(ctx *models.Context) error {
	var neededArgCount int = 0
	for range ctx.CurrentCommand.Args {
		neededArgCount++
	}
	fmt.Println(neededArgCount)
	fmt.Println(ctx.ArgsRaw)
	var args []string = strings.Fields(ctx.ArgsRaw)
	if neededArgCount > len(args) {
		return errors.New("Command is missing args, there are **" + strconv.Itoa(neededArgCount) + "** required args but you gave **" + strconv.Itoa(len(args)) + "**.")
	}

	return nil
}

func HasArgs(args []string) bool {
	if len(args) > 0 {
		return true
	} else {
		return false
	}
}
