package registry

import (
	"guillermotinoco.com/pokemonAPI/interface/controller"
	ip "guillermotinoco.com/pokemonAPI/interface/presenter"
	ir "guillermotinoco.com/pokemonAPI/interface/repository"
	"guillermotinoco.com/pokemonAPI/usecase/interactor"
	up "guillermotinoco.com/pokemonAPI/usecase/presenter"
	ur "guillermotinoco.com/pokemonAPI/usecase/repository"
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
