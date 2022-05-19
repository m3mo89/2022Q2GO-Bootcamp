package registry

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/interface/controller"

	"github.com/m3mo89/2022Q2GO-Bootcamp/infrastructure/datastore"
)

type registry struct {
	db datastore.Database
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db datastore.Database) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
