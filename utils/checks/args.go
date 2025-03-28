package checks

import (
	"errors"
)

func HasArgs(neededArgs int, args []string) error {

	return errors.New("Command is missing args.")
}
