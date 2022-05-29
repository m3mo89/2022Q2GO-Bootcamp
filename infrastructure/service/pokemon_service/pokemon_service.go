package pokemon_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
)

const apiUrl = "https://pokeapi.co/api/v2/"

type PokemonService interface {
	FindRemoteById(id int) (*model.Pokemon, error)
}

type pokemonService struct{}

func NewPokemonService() PokemonService {
	return &pokemonService{}
}

func (*pokemonService) FindRemoteById(id int) (*model.Pokemon, error) {

	var pokemon *model.Pokemon
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
