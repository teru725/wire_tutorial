// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"golab/wiretutorial"
)

// Injectors from injector.go:

func InitializeEvent(phrase string) (wiretutorial.Event, error) {
	message := wiretutorial.NewMessage(phrase)
	greeter := wiretutorial.NewGreeter(message)
	event, err := wiretutorial.NewEvent(greeter)
	if err != nil {
		return wiretutorial.Event{}, err
	}
	return event, nil
}
