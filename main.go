package main

import (
	_ "github.com/manujurado1/SportsBar-IV/docs"
	"github.com/manujurado1/SportsBar-IV/pkg/api"
)

// @title SportsBar API REST
// @version 1.0
// @description Una API para gestionar partido igualados entre grupos de amigos
func main() {
	r := api.SetupRouter()

	_ = r.Run()
}
