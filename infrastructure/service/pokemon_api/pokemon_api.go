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

func GetPokemon(id int) (*model.PokemonApi, error) {

	var pokemon *model.PokemonApi
	pokemonId := strconv.Itoa(id)

	endpoint := "pokemon/" + pokemonId

	resp, err := http.DefaultClient.Get(apiUrl + endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to reach [%v]: %v", apiUrl, err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %v", err)
	}

	json.Unmarshal(content, &pokemon)

	return pokemon, nil
}
