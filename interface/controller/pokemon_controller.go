package controller

import (
	"net/http"
	"strconv"

	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
	"github.com/m3mo89/2022Q2GO-Bootcamp/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(c Context) error
	GetPokemonById(c Context) error
}

func NewPokemonController(pk interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pk}
}

func (pc *pokemonController) GetPokemons(c Context) error {
	var p []*model.Pokemon

	p, err := pc.pokemonInteractor.Get()
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

	var p *model.Pokemon

	p, err := pc.pokemonInteractor.GetById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
