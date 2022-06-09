package service

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"
	"github.com/m3mo89/2022Q2GO-Bootcamp/test/testdata"
)

type mockDatasource struct {
	mock.Mock
}

func (datasource *mockDatasource) FindAll() ([]*entity.Pokemon, error) {
	log.Printf("Data Source Mock: Find all the pokemons")
	arg := datasource.Called()
	return arg.Get(0).([]*entity.Pokemon), arg.Error(1)
}

func (datasource *mockDatasource) FindById(id int) (*entity.Pokemon, error) {
	log.Printf("Data Source Mock: Find pokemon with id %d", id)
	arg := datasource.Called(id)
	return arg.Get(0).(*entity.Pokemon), arg.Error(1)
}

func (datasource *mockDatasource) Save(pokemon *entity.Pokemon) (*entity.Pokemon, error) {
	log.Printf("Data Source Mock: Save Pokemon %+v", pokemon)
	arg := datasource.Called(pokemon)
	return arg.Get(0).(*entity.Pokemon), arg.Error(0)
}

func TestPokemonService_GetById(t *testing.T) {
	var testCases = []struct {
		name               string
		id                 int
		response           *entity.Pokemon
		err                error
		repositoryLayer    string
		repositoryResponse *entity.Pokemon
		repositoryError    error
	}{
		{
			"Should return 1 pokemon by id from mock data source",
			6,
			&entity.Pokemon{Id: 6, Name: "charizard", Height: 17, IsDefault: true, Order: 7, Weight: 905, BaseExperience: 267, LocationAreaEncounters: "https://pokeapi.co/api/v2/pokemon/6/encounters"},
			nil,
			"mock",
			testdata.Pokemons[6],
			nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			var service PokemonService

			switch testCase.repositoryLayer {
			case "mock":
				local := &mockDatasource{}
				remote := &mockDatasource{}
				local.On("FindById", testCase.id).Return(testCase.repositoryResponse, testCase.repositoryError)
				service = NewPokemonService(local, remote)
			default:
				t.Fatalf("Should use valid data source: %v", testCase.repositoryLayer)
			}

			pokemon, err := service.GetById(testCase.id)
			t.Logf("Pokemon found: %v", pokemon)

			assert.Equal(t, testCase.response, pokemon)
			assert.Equal(t, testCase.err, err)
		})
	}
}
