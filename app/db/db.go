package db

import (
	rethink "github.com/christopherhesse/rethinkgo"
	"github.com/go-martini/martini"
)

func DB() martini.Handler {
	return func(c martini.Context) {
		session, err := rethink.Connect("localhost:28015", "countify")
		if err != nil {
			panic(err)
		}
		c.Map(session)
		defer session.Close()
		c.Next()
	}
}
