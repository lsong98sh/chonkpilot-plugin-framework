package bus

import "errors"

var (
	errNilHandler = errors.New("bus: nil handler")
	errClosed     = errors.New("bus: bus closed")
)
