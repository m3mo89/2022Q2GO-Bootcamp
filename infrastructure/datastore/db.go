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
	db := database{path: path}
	db.readData()

	return &db
}

type database struct {
	path    string
	data    []*model.Pokemon
	dataMap map[int]*model.Pokemon
}

func (d *database) readData() error {
	var records []*model.Pokemon

	defer d.convertDataToMap()

	file, err := os.Open(d.path)

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

func (d *database) convertDataToMap() {

	pokemons := make(map[int]*model.Pokemon)

	for _, value := range d.data {
		pokemons[value.Id] = value
	}

	d.dataMap = pokemons
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
