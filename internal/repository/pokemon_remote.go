package repository

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

type pokemonRemote struct{}

func NewPokemonRemote() Datasource {
	return &pokemonRemote{}
}

func (*pokemonRemote) FindAll() ([]*entity.Pokemon, error) {
	return nil, errors.New("FindAll method is not supported")
}

func (*pokemonRemote) Save(pokemon *entity.Pokemon) (*entity.Pokemon, error) {
	return nil, errors.New("Save method is not supported")
}

func (*pokemonRemote) FindById(id int) (*entity.Pokemon, error) {

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
