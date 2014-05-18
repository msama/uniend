package main

import (
	"github.com/go-martini/martini"
	"github.com/msama/uniend/admin/api"
)

func main() {
	m := martini.Classic()

	m.Group("/app", func(r martini.Router) {
		r.Get("/", api.GetApps)
		r.Get("/:id", api.GetApp)
		r.Post("/new", api.NewApp)
		r.Put("/update/:id", api.UpdateApp)
		r.Delete("/delete/:id", api.DeleteApp)
	})

	m.Run()
}
