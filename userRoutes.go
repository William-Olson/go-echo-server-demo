package main

import (
	"github.com/labstack/echo"
)

// ------------ user routes --------------

// definition
type userRoutes struct {
	route
}

// mappings
func (ur userRoutes) init() {
	ur.get("/:id/details", ur.getRoot)
	ur.get("/:id", ur.userByIdRoute)
	ur.post("/", ur.createUser)
}

// handlers
func (u userRoutes) getRoot(c echo.Context) error {
	str := "details for user " + c.Param("id")
	return c.String(200, str)
}

func (u userRoutes) userByIdRoute(c echo.Context) error {
	return c.String(200, "fetching user with id: "+c.Param("id"))
}

func (u userRoutes) createUser(c echo.Context) error {

	first := c.FormValue("first")
	last := c.FormValue("last")

	users.create(first, last)

	return c.String(200, "Ok")
}
