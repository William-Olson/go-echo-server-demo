package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	client := DB{}
	client.connect()

	// sync schema
	client.db.AutoMigrate(&User{})

	// test data creation
	client.users.create("admin", "admin")

	mountRoutes(rootRoutes{route{"/", e, &client}})
	mountRoutes(userRoutes{route{"/users", e, &client}})

	e.Logger.Fatal(e.Start(":7447"))
	defer client.db.Close()
}
