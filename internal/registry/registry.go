package registry

import (
	dddcore "cypt/internal/dddcore/adapter"
)

type registry struct {
}

type Registry interface {
	NewAppController() AppController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() AppController {
	eventBus := dddcore.NewWatermillEventBus()

	return AppController{
		User: r.NewUserController(&eventBus),
	}
}
