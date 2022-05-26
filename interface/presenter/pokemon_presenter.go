package presenter

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
)

type pokemonPresenter struct {
}

type PokemonPresenter interface {
	ResponsePokemons(p []*model.Pokemon) []*model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemons(pk []*model.Pokemon) []*model.Pokemon {
	for _, p := range pk {
		p.Name = "Mr." + p.Name
	}
	return pk
}
