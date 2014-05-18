package main

import (
	"flag"
	"github.com/go-martini/martini"
	"github.com/msama/uniend/admin/api"
	"github.com/msama/uniend/common"
	"log"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config",
		"dev-properties.json", "Configuration file")
}

func main() {
	flag.Parse()
	log.Println("Using config file", configFile)
	config, err := common.NewConfig(configFile)
	if err != nil {
		log.Fatal("Cannot read config file: ", err)
	}

	dynamo, err := common.GetServer(config)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
	tables, err := dynamo.ListTables()
	if err != nil {
		log.Fatal("Cannot fetch tables: ", err)
	}
	log.Println("Found: ", tables)

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
