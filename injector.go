//go:build wireinject
// +build wireinject

package main

import (
	"golab/wiretutorial"

	"github.com/google/wire"
)

func InitializeEvent(text string) (wiretutorial.Event, error) {
	wire.Build(wiretutorial.Set)
	return wiretutorial.Event{}, nil // return value not used - only for avoid type error
}
