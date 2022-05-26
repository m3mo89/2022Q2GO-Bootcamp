package registry

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
	"github.com/m3mo89/2022Q2GO-Bootcamp/interface/controller"
	ip "github.com/m3mo89/2022Q2GO-Bootcamp/interface/presenter"
	ir "github.com/m3mo89/2022Q2GO-Bootcamp/interface/repository"
	"github.com/m3mo89/2022Q2GO-Bootcamp/usecase/interactor"
)

type PokemonRepository interface {
	FindAll() ([]*model.Pokemon, error)
	FindById(id int) (*model.Pokemon, error)
}

type PokemonPresenter interface {
	ResponsePokemons(p []*model.Pokemon) []*model.Pokemon
}

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() PokemonRepository {
	return ir.NewPokemonRepository(r.db)
}

func (r *registry) NewPokemonPresenter() PokemonPresenter {
	return ip.NewPokemonPresenter()
}
