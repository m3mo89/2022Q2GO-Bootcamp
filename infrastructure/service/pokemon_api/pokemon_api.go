package pokemon_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
	"github.com/m3mo89/2022Q2GO-Bootcamp/infrastructure/service"
)

const apiUrl = "https://pokeapi.co/api/v2/"

type pokemonService struct{}

func NewPokemonService() service.Datasource {
	return &pokemonService{}
}

func (*pokemonService) FindAll() ([]*model.Pokemon, error) {
	return nil, errors.New("FindAll method is not supported")
}

func (*pokemonService) Save(pokemon *model.Pokemon) (*model.Pokemon, error) {
	return nil, errors.New("Save method is not supported")
}

func (*pokemonService) FindById(id int) (*model.Pokemon, error) {

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
