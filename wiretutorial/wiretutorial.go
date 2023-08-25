package wiretutorial

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewMessage,
	NewGreeter,
	NewEvent,
)
