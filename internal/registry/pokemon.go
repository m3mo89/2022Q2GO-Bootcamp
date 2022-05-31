package registry

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/controller"
	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"
	ir "github.com/m3mo89/2022Q2GO-Bootcamp/internal/repository"
)

type PokemonRepository interface {
	FindAll() ([]*entity.Pokemon, error)
	FindById(id int) (*entity.Pokemon, error)
	Save(pokemon *entity.Pokemon) (*entity.Pokemon, error)
}

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonRepository())
}

func (r *registry) NewPokemonRepository() PokemonRepository {
	return ir.NewPokemonRepository(r.local, r.remote)
}
