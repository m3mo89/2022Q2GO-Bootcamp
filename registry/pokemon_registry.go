package registry

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/interface/controller"
	ip "github.com/m3mo89/2022Q2GO-Bootcamp/interface/presenter"
	ir "github.com/m3mo89/2022Q2GO-Bootcamp/interface/repository"
	"github.com/m3mo89/2022Q2GO-Bootcamp/usecase/interactor"
	up "github.com/m3mo89/2022Q2GO-Bootcamp/usecase/presenter"
	ur "github.com/m3mo89/2022Q2GO-Bootcamp/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository(r.db)
}

func (r *registry) NewPokemonPresenter() up.PokemonPresenter {
	return ip.NewPokemonPresenter()
}
