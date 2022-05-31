package registry

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/controller"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/repository"
)

type registry struct {
	local  repository.Datasource
	remote repository.Datasource
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(local repository.Datasource, remote repository.Datasource) Registry {
	return &registry{local, remote}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
