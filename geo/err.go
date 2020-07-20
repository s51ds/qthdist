package geo

import (
	"errors"
	"fmt"
)

func illegalLocatorError(arg string) error {
	return errors.New(fmt.Sprintf("Illegal argumet value! qthLocator=%s", arg))
}
