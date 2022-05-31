package repository

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
)

type PokemonRepository interface {
	FindAll() ([]*model.Pokemon, error)
	FindById(id int) (*model.Pokemon, error)
	Save(pokemon *model.Pokemon) (*model.Pokemon, error)
}

type Datasource interface {
	FindAll() ([]*model.Pokemon, error)
	FindById(id int) (*model.Pokemon, error)
	Save(pokemon *model.Pokemon) (*model.Pokemon, error)
}

type pokemonRepository struct {
	srcLocal  Datasource
	srcRemote Datasource
}

func NewPokemonRepository(local, remote Datasource) PokemonRepository {
	return &pokemonRepository{local, remote}
}

func (pr *pokemonRepository) FindAll() ([]*model.Pokemon, error) {
	pokemons, err := pr.srcLocal.FindAll()

	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (pr *pokemonRepository) FindById(id int) (*model.Pokemon, error) {
	pokemon, err := pr.srcLocal.FindById(id)

	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

func (pr *pokemonRepository) Save(pokemon *model.Pokemon) (*model.Pokemon, error) {
	pokemon, err := pr.srcLocal.Save(pokemon)

	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
