package main

import (
	"github.com/labstack/echo"
)

// ------------------------------------------------------

func main() {
	e := echo.New()

	mountRoutes(rootRoutes{route{"/", e}})
	mountRoutes(userRoutes{route{"/users", e}})
	// users := userRoutes{route{"/users", e}}
	// mountRoutes(users)

	e.Logger.Fatal(e.Start(":7447"))
}
