package mock

import (
	"log"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"
	"github.com/stretchr/testify/mock"
)

type mockPokemonService struct {
	mock.Mock
}

func NewPokemonServiceMock() *mockPokemonService {
	return &mockPokemonService{}
}

func (service *mockPokemonService) GetAll() ([]*entity.Pokemon, error) {
	log.Printf("Pokemon Service Mock: Get all the pokemons")
	arg := service.Called()

	if arg.Get(0) != nil {
		return arg.Get(0).([]*entity.Pokemon), arg.Error(1)
	}
	return nil, arg.Error(1)
}

func (service *mockPokemonService) GetById(id int) (*entity.Pokemon, error) {
	log.Printf("Pokemon Service Mock: Get pokemon by id %d", id)
	arg := service.Called(id)

	if arg.Get(0) != nil {
		return arg.Get(0).(*entity.Pokemon), arg.Error(1)
	}
	return nil, arg.Error(1)
}

func (service *mockPokemonService) GetAllWithWorker(item_type string, items, items_per_workers int) ([]*entity.Pokemon, error) {
	log.Printf("Pokemon Service Mock: Get all the pokemons usig worker pool %s %d %d", item_type, items, items_per_workers)
	arg := service.Called(item_type, items, items_per_workers)

	if arg.Get(0) != nil {
		return arg.Get(0).([]*entity.Pokemon), arg.Error(1)
	}
	return nil, arg.Error(1)
}
