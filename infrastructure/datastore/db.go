package datastore

import (
	"errors"
	"log"
	"os"

	"github.com/gocarina/gocsv"

	"github.com/m3mo89/2022Q2GO-Bootcamp/domain/model"
)

type Database interface {
	FindAll() ([]*model.Pokemon, error)
	FindById(id int) (*model.Pokemon, error)
}

func NewDatabase(path string) Database {
	records, _ := readData(path)
	recordsMap := convertDataToMap(records)

	return &database{path, records, recordsMap}
}

type database struct {
	path    string
	data    []*model.Pokemon
	dataMap map[int]*model.Pokemon
}

func readData(path string) ([]*model.Pokemon, error) {
	var records []*model.Pokemon

	defer convertDataToMap(records)

	file, err := os.Open(path)

	defer func() {
		err := file.Close()

		log.Println(err)
	}()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		log.Println(err)
		return nil, err
	}

	return records, nil
}

func convertDataToMap(records []*model.Pokemon) (pokemons map[int]*model.Pokemon) {

	if pokemons == nil {
		pokemons = make(map[int]*model.Pokemon)
	}

	for _, value := range records {
		pokemons[value.Id] = value
	}

	return pokemons
}

func (d *database) FindAll() ([]*model.Pokemon, error) {
	return d.data, nil
}

func (d *database) FindById(id int) (*model.Pokemon, error) {

	pokemon, ok := d.dataMap[id]

	if !ok {
		return nil, errors.New("the pokemon was not found")
	}
	return pokemon, nil
}
