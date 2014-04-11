package controllers

import (
	"../models"
	"github.com/christopherhesse/rethinkgo"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
)

type AccountsController struct {
}

func (c AccountsController) GetAccounts(render render.Render, session *rethinkgo.Session) {
	var accounts []*models.Account
	err := rethinkgo.Table("accounts").Run(session).All(&accounts)
	if err != nil {
		panic(err)
	}
	render.JSON(200, accounts)
}

func (c AccountsController) GetAccount(params martini.Params, render render.Render, session *rethinkgo.Session) {
	var account *models.Account
	err := rethinkgo.Table("accounts").Get(params["id"]).Run(session).One(&account)
	if err != nil {
		panic(err)
	}
	render.JSON(200, account)
}

func (c AccountsController) CreateAccount(account models.Account, render render.Render, session *rethinkgo.Session) {
	log.Println("account", account)
	rethinkgo.Table("accounts").Insert(account).Run(session)
	c.GetAccounts(render, session)
	//render.JSON(200, "ok")
}
