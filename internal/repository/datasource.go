package repository

import "github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"

type Datasource interface {
	FindAll() ([]*entity.Pokemon, error)
	FindById(id int) (*entity.Pokemon, error)
	Save(pokemon *entity.Pokemon) (*entity.Pokemon, error)
	FindAllWithWorker(item_type string, items, items_per_workers int) ([]*entity.Pokemon, error)
}
