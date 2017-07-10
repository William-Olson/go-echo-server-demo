package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"time"
)

type rootRoutes struct {
	route
}

// mappings
func (r rootRoutes) mount() {

	r.group.GET("/", r.getRoot)
	r.group.POST("/login", r.login)

}

// handlers
func (r rootRoutes) getRoot(c echo.Context) error {

	return c.String(200, "this is the root route")

}

func (r rootRoutes) login(c echo.Context) error {

	un := c.FormValue("username")
	pw := c.FormValue("password")

	user, err := r.db.Users.Login(un, pw)

	if err != nil {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.First + " " + user.Last
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{
		"token": t,
		"ok":    "true",
	})

}
