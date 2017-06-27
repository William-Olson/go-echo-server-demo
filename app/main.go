package main

import (
	"api"
	"models"
)

func main() {

	db := models.DB{}
	sv := api.Server{Db: &db}

	// setup
	db.Start()
	sv.Start()

	defer db.Client.Close()

}
