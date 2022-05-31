package registry

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/controller"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/service"
)

type registry struct {
	local  service.Datasource
	remote service.Datasource
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(local service.Datasource, remote service.Datasource) Registry {
	return &registry{local, remote}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
