package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"strconv"
)

type userRoutes struct {
	route
}

// mappings
func (ur userRoutes) mount() {

	// use auth middleware
	ur.group.Use(middleware.JWT([]byte(jwtSecret)))

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

	username := c.FormValue("username")
	first := c.FormValue("first")
	last := c.FormValue("last")
	pw := c.FormValue("password")

	u.db.Users.Create(username, first, last, pw)

	return c.String(200, "Ok")

}
