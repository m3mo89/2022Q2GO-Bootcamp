package interactor

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
	"github.com/m3mo89/2022Q2GO-Bootcamp/usecase/presenter"
	"github.com/m3mo89/2022Q2GO-Bootcamp/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	Get() ([]*model.Pokemon, error)
	GetById(id int) (*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (pk *pokemonInteractor) Get() ([]*model.Pokemon, error) {
	p, err := pk.PokemonRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return pk.PokemonPresenter.ResponsePokemons(p), nil
}

func (pk *pokemonInteractor) GetById(id int) (*model.Pokemon, error) {
	p, err := pk.PokemonRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return p, nil
}
