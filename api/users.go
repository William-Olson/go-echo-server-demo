package api

import (
	"github.com/labstack/echo"
)

type userRoutes struct {
	route
}

// mappings
func (ur userRoutes) mount() {

	ur.group.GET("/:id/details", ur.getRoot)
	ur.group.GET("/:id", ur.userByIdRoute)
	ur.group.POST("/", ur.createUser)

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

	u.db.Users.Create(first, last)

	return c.String(200, "Ok")

}
