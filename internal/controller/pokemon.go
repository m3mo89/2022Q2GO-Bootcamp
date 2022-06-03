package controller

import (
	"net/http"
	"strconv"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"
)

type pokemonController struct {
	pokemonService PokemonService
}

type PokemonService interface {
	GetAll() ([]*entity.Pokemon, error)
	GetById(id int) (*entity.Pokemon, error)
}

type PokemonController interface {
	GetPokemons(c Context) error
	GetPokemonById(c Context) error
}

func NewPokemonController(service PokemonService) PokemonController {
	return &pokemonController{service}
}

func (pc *pokemonController) GetPokemons(c Context) error {
	var p []*entity.Pokemon

	p, err := pc.pokemonService.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func (pc *pokemonController) GetPokemonById(c Context) error {

	id, errAtoi := strconv.Atoi(c.Param("id"))

	if errAtoi != nil {
		return errAtoi
	}

	var p *entity.Pokemon

	p, err := pc.pokemonService.GetById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
