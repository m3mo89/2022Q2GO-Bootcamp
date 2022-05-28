package pokemon_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
)

const apiUrl = "https://pokeapi.co/api/v2/"

type PokemonApi interface {
	GetPokemon(id int) (*model.PokemonApi, error)
}

type pokemonApi struct{}

func NewPokemonApi() PokemonApi {
	return &pokemonApi{}
}

func (*pokemonApi) GetPokemon(id int) (*model.PokemonApi, error) {

	var pokemon *model.PokemonApi
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
