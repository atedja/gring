package gring

import (
	"errors"
)

var ErrInvalidOperationOnDetachedNode = errors.New("Cannot do the specified operation on detached nodes.")
var ErrEmptyRing = errors.New("Cannot do the specified operation on an empty ring.")
