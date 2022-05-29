package registry

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/interface/controller"

	"github.com/m3mo89/2022Q2GO-Bootcamp/infrastructure/datastore"
	"github.com/m3mo89/2022Q2GO-Bootcamp/infrastructure/service/pokemon_service"
)

type registry struct {
	db      datastore.Database
	service pokemon_service.PokemonService
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db datastore.Database, service pokemon_service.PokemonService) Registry {
	return &registry{db, service}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
