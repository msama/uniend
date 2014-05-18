package api

import (
	"fmt"
	"github.com/go-martini/martini"
)

func GetApps(params martini.Params) string {
	return fmt.Sprintf("Hello %s", params["name"])
}

func GetApp(params martini.Params) string {
	return fmt.Sprintf("Hello %s", params["name"])
}
func NewApp(params martini.Params) string {
	return fmt.Sprintf("Hello %s", params["name"])
}
func UpdateApp(params martini.Params) string {
	return fmt.Sprintf("Hello %s", params["name"])
}
func DeleteApp(params martini.Params) string {
	return fmt.Sprintf("Hello %s", params["name"])
}
