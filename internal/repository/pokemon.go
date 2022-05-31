package repository

import (
	"errors"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"
)

type PokemonRepository interface {
	FindAll() ([]*entity.Pokemon, error)
	FindById(id int) (*entity.Pokemon, error)
	Save(pokemon *entity.Pokemon) (*entity.Pokemon, error)
}

type Datasource interface {
	FindAll() ([]*entity.Pokemon, error)
	FindById(id int) (*entity.Pokemon, error)
	Save(pokemon *entity.Pokemon) (*entity.Pokemon, error)
}

type pokemonRepository struct {
	srcLocal  Datasource
	srcRemote Datasource
}

func NewPokemonRepository(local, remote Datasource) PokemonRepository {
	return &pokemonRepository{local, remote}
}

func (pr *pokemonRepository) FindAll() ([]*entity.Pokemon, error) {
	pokemons, err := pr.srcLocal.FindAll()

	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (pr *pokemonRepository) FindById(id int) (*entity.Pokemon, error) {
	pokemon, err := pr.srcLocal.FindById(id)

	if pokemon != nil && err != nil {
		return nil, err
	}

	if pokemon == nil {
		pokemon, err = pr.srcRemote.FindById(id)

		if err != nil {
			return nil, err
		}

		if pokemon == nil {
			return nil, errors.New("pokemon not found")
		}

		pokemon, err = pr.srcLocal.Save(pokemon)

		if err != nil {
			return nil, err
		}
	}

	return pokemon, nil
}

func (pr *pokemonRepository) Save(pokemon *entity.Pokemon) (*entity.Pokemon, error) {
	pokemon, err := pr.srcLocal.Save(pokemon)

	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
