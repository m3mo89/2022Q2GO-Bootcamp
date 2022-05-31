package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"

	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/registry"
	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/router"
	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/service"
)

func main() {
	fileName := "data/pokemon.csv"

	local := service.NewDatabase(fileName)

	remote := service.NewPokemonService()

	r := registry.NewRegistry(local, remote)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost:8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
