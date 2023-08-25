package wiretutorial

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(Message), "*"),
	NewGreeter,
	NewEvent,
)
