package main

import (
	"github.com/go-martini/martini"
	"github.com/msama/uniend/frontend/api"
)

func main() {
	m := martini.Classic()

	m.Group("/storage", func(r martini.Router) {
		r.Get("/", api.GetKeys)
		r.Get("/:key", api.GetValue)
		r.Post("/", api.PostValue)
		r.Put("/:key", api.PutValue)
		r.Delete(":key", api.DeleteValue)
	})

	m.Run()
}
