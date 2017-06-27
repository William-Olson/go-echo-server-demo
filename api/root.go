package api

import (
	"github.com/labstack/echo"
)

type rootRoutes struct {
	route
}

// mappings
func (r rootRoutes) mount() {

	r.group.GET("/", r.getRoot)

}

// handlers
func (r rootRoutes) getRoot(c echo.Context) error {

	return c.String(200, "this is the root route")

}
