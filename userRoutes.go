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
}

// handlers
func (u userRoutes) getRoot(c echo.Context) error {
	str := "details for user " + c.Param("id")
	return c.String(200, str)
}

func (u userRoutes) userByIdRoute(c echo.Context) error {
	return c.String(200, "fetching user with id: "+c.Param("id"))
}
