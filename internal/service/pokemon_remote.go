package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"
)

const apiUrl = "https://pokeapi.co/api/v2/"

type pokemonService struct{}

func NewPokemonService() Datasource {
	return &pokemonService{}
}

func (*pokemonService) FindAll() ([]*entity.Pokemon, error) {
	return nil, errors.New("FindAll method is not supported")
}

func (*pokemonService) Save(pokemon *entity.Pokemon) (*entity.Pokemon, error) {
	return nil, errors.New("Save method is not supported")
}

func (*pokemonService) FindById(id int) (*entity.Pokemon, error) {

	var pokemon *entity.Pokemon
	pokemonId := strconv.Itoa(id)

	endpoint := "pokemon/" + pokemonId

	response, err := http.DefaultClient.Get(apiUrl + endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to reach [%v]: %v", apiUrl, err)
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %v", err)
	}

	json.Unmarshal(content, &pokemon)

	return pokemon, nil
}
