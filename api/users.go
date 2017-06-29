package api

import (
	"github.com/labstack/echo"
	"strconv"
)

type userRoutes struct {
	route
}

// mappings
func (ur userRoutes) mount() {

	ur.group.GET("/:id", ur.getUser)
	ur.group.GET("/", ur.getAll)
	ur.group.POST("/", ur.createUser)

}

// handlers
func (u userRoutes) getUser(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		return err
	}

	user := u.db.Users.ById(uint(id))

	return c.JSON(200, user)

}

func (u userRoutes) getAll(c echo.Context) error {

	return c.JSON(200, u.db.Users.GetAll())

}

func (u userRoutes) createUser(c echo.Context) error {

	first := c.FormValue("first")
	last := c.FormValue("last")

	u.db.Users.Create(first, last)

	return c.String(200, "Ok")

}
