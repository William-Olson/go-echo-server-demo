package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db := DB{}

	// db setup
	db.connect()
	db.sync()
	db.addTestData()

	// route setup
	mountRoutes(rootRoutes{route{"/", e, &db}})
	mountRoutes(userRoutes{route{"/users", e, &db}})

	e.Logger.Fatal(e.Start(":7447"))
	defer db.client.Close()
}
