package registry

type registry struct {
}

type Registry interface {
	NewAppController() AppController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() AppController {
	return AppController{
		User: r.NewUserController(),
	}
}
