package presenter

import "github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"

type PokemonPresenter interface {
	ResponsePokemons(p []*model.Pokemon) []*model.Pokemon
}
