package repository

import "github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"

type PokemonRepository interface {
	FindAll() ([]*model.Pokemon, error)
	FindById(id int) (*model.Pokemon, error)
}
