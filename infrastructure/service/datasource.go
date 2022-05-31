package service

import "github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"

type Datasource interface {
	FindAll() ([]*model.Pokemon, error)
	FindById(id int) (*model.Pokemon, error)
	Save(pokemon *model.Pokemon) (*model.Pokemon, error)
}
