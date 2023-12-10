package event

import (
	"github.com/fwidjaya20/symphonic/contracts/foundation"
	"github.com/fwidjaya20/symphonic/facades"
)

type ServiceProvider struct{}

func (sp *ServiceProvider) Boot(_ foundation.Application) {}

func (sp *ServiceProvider) Register(_ foundation.Application) {
	kernel := Kernel{}

	facades.Event().Register(kernel.Subscribers())
}
