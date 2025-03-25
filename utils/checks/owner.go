package checks

import (
	"errors"
	"gobot/models"
	"strconv"
)

func IsOwner(ctx *models.Context) error {
	var flag bool = true
	for ownerID := range ctx.Client.Owners {
		if strconv.Itoa(ownerID) == ctx.Author.ID {
			flag = false
		}
	}
	if flag {
		return errors.New("Owner only command.")
	}
	return nil
}
