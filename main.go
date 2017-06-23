package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var database *gorm.DB

func main() {
	e := echo.New()

	client := DB{}
	database = client.connect()
	client.db = database

	// sync schema
	client.db.AutoMigrate(&User{})

	// test data creation
	users.create("admin", "admin")

	mountRoutes(rootRoutes{route{"/", e, &client}})
	mountRoutes(userRoutes{route{"/users", e, &client}})

	e.Logger.Fatal(e.Start(":7447"))
	defer database.Close()
}
