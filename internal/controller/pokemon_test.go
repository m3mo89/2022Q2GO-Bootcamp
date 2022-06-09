package controller

import (
	"errors"
	"net/http"
	"testing"

	customMock "github.com/m3mo89/2022Q2GO-Bootcamp/test/mock"
	"github.com/m3mo89/2022Q2GO-Bootcamp/test/testdata"

	"github.com/stretchr/testify/assert"
)

func TestPokemonController_GetPokemons(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (PokemonController, Context)
		assert  func(t *testing.T, context Context, err error)
	}{
		"success": {
			arrange: func(t *testing.T) (PokemonController, Context) {
				mockService := customMock.NewPokemonServiceMock()
				mockService.On("GetAll").Return(testdata.Pokemons, nil)
				controller := NewPokemonController(mockService)
				mockContext := customMock.NewContextMock()
				mockContext.On("JSON", http.StatusOK, testdata.Pokemons).Return(nil)
				return controller, mockContext
			},
			assert: func(t *testing.T, context Context, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusOK, 200)
			},
		},
		"error": {
			arrange: func(t *testing.T) (PokemonController, Context) {
				mockService := customMock.NewPokemonServiceMock()
				mockService.On("GetAll").Return(nil, errors.New("fake error"))
				controller := NewPokemonController(mockService)
				context := customMock.NewContextMock()
				return controller, context
			},
			assert: func(t *testing.T, context Context, err error) {
				assert.NotNil(t, err)
				assert.NotNil(t, context)
				assert.Nil(t, nil)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			controller, context := tt.arrange(t)
			err := controller.GetPokemons(context)
			tt.assert(t, context, err)
		})
	}
}

func TestPokemonController_GetPokemonById(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (PokemonController, Context)
		assert  func(t *testing.T, context Context, err error)
	}{
		"success": {
			arrange: func(t *testing.T) (PokemonController, Context) {
				mockService := customMock.NewPokemonServiceMock()
				mockService.On("GetById", 7).Return(testdata.Pokemons[0], nil)
				controller := NewPokemonController(mockService)

				mockContext := customMock.NewContextMock()
				mockContext.On("Param", "id").Return("7")
				mockContext.On("JSON", http.StatusOK, testdata.Pokemons[0]).Return(nil)
				return controller, mockContext
			},
			assert: func(t *testing.T, context Context, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusOK, 200)
			},
		},
		"error": {
			arrange: func(t *testing.T) (PokemonController, Context) {
				mockService := customMock.NewPokemonServiceMock()
				mockService.On("GetById", 40).Return(nil, errors.New("fake error"))
				controller := NewPokemonController(mockService)

				mockContext := customMock.NewContextMock()
				mockContext.On("Param", "id").Return("40")
				mockContext.On("JSON", http.StatusOK, nil).Return(nil)
				return controller, mockContext
			},
			assert: func(t *testing.T, context Context, err error) {
				assert.NotNil(t, err)
				assert.NotNil(t, context)
				assert.Equal(t, http.StatusOK, 200)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			controller, context := tt.arrange(t)
			err := controller.GetPokemonById(context)
			tt.assert(t, context, err)
		})
	}
}
