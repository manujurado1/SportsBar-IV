package main

import "github.com/manujurado1/SportsBar-IV/pkg/api"

func main() {
	r := api.SetupRouter()

	_ = r.Run()
}
