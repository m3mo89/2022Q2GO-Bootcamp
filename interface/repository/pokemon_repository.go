package repository

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
)

type PokemonRepository interface {
	FindAll() ([]*model.Pokemon, error)
	FindById(id int) (*model.Pokemon, error)
	FindRemoteById(id int) (*model.RemotePokemon, error)
}

type PokemonService interface {
	GetRemotePokemon(id int) (*model.RemotePokemon, error)
}

type Database interface {
	FindAll() ([]*model.Pokemon, error)
	FindById(id int) (*model.Pokemon, error)
}

type pokemonRepository struct {
	db      Database
	service PokemonService
}

func NewPokemonRepository(db Database, service PokemonService) PokemonRepository {
	return &pokemonRepository{db, service}
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

func (pr *pokemonRepository) FindRemoteById(id int) (*model.RemotePokemon, error) {
	pokemon, err := pr.service.GetRemotePokemon(id)

	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
