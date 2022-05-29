package interactor

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
)

type pokemonInteractor struct {
	PokemonRepository PokemonRepository
	PokemonPresenter  PokemonPresenter
}

type PokemonRepository interface {
	FindAll() ([]*model.Pokemon, error)
	FindById(id int) (*model.Pokemon, error)
	FindRemoteById(id int) (*model.RemotePokemon, error)
	Save() (*model.Pokemon, error)
}

type PokemonPresenter interface {
	ResponsePokemons(p []*model.Pokemon) []*model.Pokemon
}

type PokemonInteractor interface {
	Get() ([]*model.Pokemon, error)
	GetById(id int) (*model.Pokemon, error)
	GetRemoteById(id int) (*model.RemotePokemon, error)
	Save() (*model.Pokemon, error)
}

func NewPokemonInteractor(r PokemonRepository, p PokemonPresenter) PokemonInteractor {
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

func (pk *pokemonInteractor) GetRemoteById(id int) (*model.RemotePokemon, error) {
	p, err := pk.PokemonRepository.FindRemoteById(id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pk *pokemonInteractor) Save() (*model.Pokemon, error) {
	p, err := pk.PokemonRepository.Save()
	if err != nil {
		return nil, err
	}

	return p, nil
}
