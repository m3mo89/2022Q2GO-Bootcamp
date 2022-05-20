package repository

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
	"github.com/m3mo89/2022Q2GO-Bootcamp/infrastructure/datastore"
	"github.com/m3mo89/2022Q2GO-Bootcamp/usecase/repository"
)

type pokemonRepository struct {
	db datastore.Database
}

func NewPokemonRepository(db datastore.Database) repository.PokemonRepository {
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
