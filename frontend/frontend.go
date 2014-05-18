package main

import (
	"flag"
	"github.com/crowdmob/goamz/dynamodb"
	"github.com/go-martini/martini"
	"github.com/msama/uniend/common"
	"github.com/msama/uniend/frontend/api"
	"log"
)

var configFile string
var dynamo dynamodb.Server

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

	m.Group("/storage", func(r martini.Router) {
		r.Get("/", api.GetKeys)
		r.Get("/:key", api.GetValue)
		r.Post("/", api.PostValue)
		r.Put("/:key", api.PutValue)
		r.Delete(":key", api.DeleteValue)
	})

	m.Run()
}
