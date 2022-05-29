package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"

	"github.com/m3mo89/2022Q2GO-Bootcamp/infrastructure/datastore"
	"github.com/m3mo89/2022Q2GO-Bootcamp/infrastructure/router"
	"github.com/m3mo89/2022Q2GO-Bootcamp/infrastructure/service/pokemon_service"
	"github.com/m3mo89/2022Q2GO-Bootcamp/registry"
)

func main() {
	fileName := "infrastructure/data/pokemon.csv"

	db := datastore.NewDatabase(fileName)

	service := pokemon_service.NewPokemonService()

	r := registry.NewRegistry(db, service)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost:8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
