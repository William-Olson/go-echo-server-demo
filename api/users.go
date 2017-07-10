package api

import (
	"github.com/labstack/echo"
	"utils"
)

type userRoutes struct {
	route
}

/*

	Map users routes

*/
func (ur userRoutes) mount() {

	ur.group.GET("/:id", ur.getUser)
	ur.group.GET("/", ur.getAll)
	ur.group.POST("/", ur.createUser)

}

/*

	Fetch a user by id

*/
func (u userRoutes) getUser(c echo.Context) error {

	id, err := utils.ConvertId(c.Param("id"))

	if err != nil {
		return err
	}

	user := u.db.Users.ById(id)

	return c.JSON(200, user)

}

/*

	Fetch all users

*/
func (u userRoutes) getAll(c echo.Context) error {

	return c.JSON(200, u.db.Users.GetAll())

}

/*

	Create a user

*/
func (u userRoutes) createUser(c echo.Context) error {

	username := c.FormValue("username")
	first := c.FormValue("first")
	last := c.FormValue("last")
	pw := c.FormValue("password")

	u.db.Users.Create(username, first, last, pw)
	resp := utils.NewResponse("ok", true)

	return c.JSON(200, resp.Payload)

}
