package gring

import (
	"errors"
)

var ErrInvalidOperationOnDetachedNode = errors.New("Cannot do the specified operation on detached nodes.")
