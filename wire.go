//go:build wireinject
// +build wireinject

package main

import (
	"golab/wiretutorial"

	"github.com/google/wire"
)

func InitializeEvent(phrase string) (wiretutorial.Event, error) {
	wire.Build(wiretutorial.NewEvent, wiretutorial.NewGreeter, wiretutorial.NewMessage)
	return wiretutorial.Event{}, nil // return value not used - only for avoid type error
}
