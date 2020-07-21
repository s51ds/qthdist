package internal

import (
	"errors"
	"fmt"
)

func IllegalLocatorError(arg string) error {
	return errors.New(fmt.Sprintf("Illegal argumet value! qthLocator=%s", arg))
}
