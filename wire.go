//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeEvent() Event {
	wire.Build(GetMessage, GetGreeter, GetEvent) //we can pass in these dependencies in any order.
	return Event{}
}
