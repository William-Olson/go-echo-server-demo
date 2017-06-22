package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	db := DB{}
	db.init()

	mountRoutes(rootRoutes{route{"/", e, &db}})
	mountRoutes(userRoutes{route{"/users", e, &db}})

	e.Logger.Fatal(e.Start(":7447"))
}
