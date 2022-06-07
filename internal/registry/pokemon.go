package registry

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/controller"
	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"
	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/service"
)

type PokemonService interface {
	GetAll() ([]*entity.Pokemon, error)
	GetById(id int) (*entity.Pokemon, error)
	GetAllWithWorker(item_type string, items, items_per_workers int) ([]*entity.Pokemon, error)
}

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonService())
}

func (r *registry) NewPokemonService() PokemonService {
	return service.NewPokemonService(r.local, r.remote)
}
