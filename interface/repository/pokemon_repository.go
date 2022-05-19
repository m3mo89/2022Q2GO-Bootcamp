package repository

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
	"github.com/m3mo89/2022Q2GO-Bootcamp/infrastructure/datastore"
)

type PokemonRepository interface {
	FindAll() ([]*model.Pokemon, error)
	FindById(id int) (*model.Pokemon, error)
}

type pokemonRepository struct {
	db datastore.Database
}

func NewPokemonRepository(db datastore.Database) PokemonRepository {
	return &pokemonRepository{db}
}

func (pr *pokemonRepository) FindAll() ([]*model.Pokemon, error) {
	pokemons, err := pr.db.FindAll()

	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (pr *pokemonRepository) FindById(id int) (*model.Pokemon, error) {
	pokemon, err := pr.db.FindById(id)

	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
