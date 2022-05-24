package presenter

import (
	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
	"github.com/m3mo89/2022Q2GO-Bootcamp/usecase/presenter"
)

type pokemonPresenter struct {
}

func NewPokemonPresenter() presenter.PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemons(pk []*model.Pokemon) []*model.Pokemon {
	for _, p := range pk {
		p.Name = "Mr." + p.Name
	}
	return pk
}
