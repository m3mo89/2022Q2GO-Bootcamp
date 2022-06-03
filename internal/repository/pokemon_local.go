package repository

import (
	"errors"
	"log"
	"os"

	"github.com/gocarina/gocsv"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/entity"
)

func NewPokemonLocal(path string) Datasource {
	db := pokemonLocal{path: path}
	db.readData()

	return &db
}

type pokemonLocal struct {
	path    string
	data    []*entity.Pokemon
	dataMap map[int]*entity.Pokemon
}

func (d *pokemonLocal) readData() error {
	var records []*entity.Pokemon

	defer d.convertDataToMap()

	file, err := os.OpenFile(d.path, os.O_RDWR|os.O_CREATE, os.ModePerm)

	defer func() {
		err := file.Close()

		log.Println(err)
	}()

	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		log.Println(err)
		return err
	}

	d.data = records

	return nil
}

func (d *pokemonLocal) writeData(pokemon *entity.Pokemon) error {
	clientsFile, err := os.OpenFile(d.path, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		log.Println(err)
	}

	defer clientsFile.Close()

	d.data = append(d.data, pokemon)

	err = gocsv.MarshalFile(&d.data, clientsFile)

	if err != nil {
		log.Println(err)
	}

	d.dataMap[int(pokemon.Id)] = pokemon

	return nil
}

func (d *pokemonLocal) convertDataToMap() {

	pokemons := make(map[int]*entity.Pokemon)

	for _, value := range d.data {
		pokemons[int(value.Id)] = value
	}

	d.dataMap = pokemons
}

func (d *pokemonLocal) FindAll() ([]*entity.Pokemon, error) {
	return d.data, nil
}

func (d *pokemonLocal) FindById(id int) (*entity.Pokemon, error) {

	pokemon, ok := d.dataMap[id]

	if !ok {
		return nil, errors.New("the pokemon was not found")
	}
	return pokemon, nil
}

func (d *pokemonLocal) Save(pokemon *entity.Pokemon) (*entity.Pokemon, error) {

	err := d.writeData(pokemon)

	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
