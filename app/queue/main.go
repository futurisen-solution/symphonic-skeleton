package main

import (
	"github.com/futurisen-solution/symphonic-skeleton/bootstrap"
	"github.com/futurisen-solution/symphonic-skeleton/bootstrap/event"
	"github.com/fwidjaya20/symphonic/facades"
)

func main() {
	bootstrap.Boot()

	facades.Event().Register(event.Kernel{}.Subscribers())

	select {}
}
