package main

import (
	"./app/controllers"
	"./app/db"
	"./app/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

  m.Use(martini.Static("app/public"))

	m.Use(db.DB())

	var accountsController *controllers.AccountsController = new(controllers.AccountsController)

	m.Group("/accounts", func(r martini.Router) {
		r.Get("", accountsController.GetAccounts)
		r.Get("/:id", accountsController.GetAccount)
		r.Post("", binding.Bind(models.Account{}), accountsController.CreateAccount)
	})

	m.Run()
}
