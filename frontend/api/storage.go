package api

import (
	"fmt"
	"github.com/go-martini/martini"
)

func GetKeys(params martini.Params) string {
	return fmt.Sprintf("GetKeys ")
}

func GetValue(params martini.Params) string {
	return fmt.Sprintf("GetValue for key %s", params["key"])
}

func PostValue(params martini.Params) string {
	return fmt.Sprintf("PostValue key value pair")
}

func PutValue(params martini.Params) string {
	return fmt.Sprintf("PutValue for key %s", params["key"])
}

func DeleteValue(params martini.Params) string {
	return fmt.Sprintf("DeleteValue for key %s", params["key"])
}
