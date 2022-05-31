package controller

import (
	"net/http"
	"strconv"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"
)

type pokemonController struct {
	pokemonRepository PokemonRepository
}

type PokemonRepository interface {
	FindAll() ([]*entity.Pokemon, error)
	FindById(id int) (*entity.Pokemon, error)
	Save(pokemon *entity.Pokemon) (*entity.Pokemon, error)
}

type PokemonController interface {
	GetPokemons(c Context) error
	GetPokemonById(c Context) error
}

func NewPokemonController(repo PokemonRepository) PokemonController {
	return &pokemonController{repo}
}

func (pc *pokemonController) GetPokemons(c Context) error {
	var p []*entity.Pokemon

	p, err := pc.pokemonRepository.FindAll()
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

	p, err := pc.pokemonRepository.FindById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
