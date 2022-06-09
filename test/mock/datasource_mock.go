package mock

import (
	"log"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"
	"github.com/stretchr/testify/mock"
)

type mockDatasource struct {
	mock.Mock
}

func NewDatasourceMock() *mockDatasource {
	return &mockDatasource{}
}

func (datasource *mockDatasource) FindAll() ([]*entity.Pokemon, error) {
	log.Printf("Data Source Mock: Find all the pokemons")
	arg := datasource.Called()

	if arg.Get(0) != nil {
		return arg.Get(0).([]*entity.Pokemon), arg.Error(1)
	}
	return nil, arg.Error(1)
}

func (datasource *mockDatasource) FindById(id int) (*entity.Pokemon, error) {
	log.Printf("Data Source Mock: Find pokemon with id %d", id)
	arg := datasource.Called(id)

	if arg.Get(0) != nil {
		return arg.Get(0).(*entity.Pokemon), arg.Error(1)
	}
	return nil, arg.Error(1)
}

func (datasource *mockDatasource) Save(pokemon *entity.Pokemon) (*entity.Pokemon, error) {
	log.Printf("Data Source Mock: Save Pokemon %+v", pokemon)
	arg := datasource.Called(pokemon)
	return arg.Get(0).(*entity.Pokemon), arg.Error(0)
}

func (datasource *mockDatasource) FindAllWithWorker(item_type string, items, items_per_workers int) ([]*entity.Pokemon, error) {
	log.Printf("Data Source Mock: Find all the pokemons usig worker pool %s %d %d", item_type, items, items_per_workers)
	arg := datasource.Called(item_type, items, items_per_workers)

	if arg.Get(0) != nil {
		return arg.Get(0).([]*entity.Pokemon), arg.Error(1)
	}
	return nil, arg.Error(1)
}
