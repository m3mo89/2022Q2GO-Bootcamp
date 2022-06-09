package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m3mo89/2022Q2GO-Bootcamp/internal/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemons", func(context echo.Context) error { return c.GetPokemons(context) })

	e.GET("/pokemonById/:id", func(context echo.Context) error { return c.GetPokemonById(context) })

	e.GET("/pokemonsWorker/:type/:items/:items_per_workers", func(context echo.Context) error { return c.GetPokemonsWithWorker(context) })

	return e
}
