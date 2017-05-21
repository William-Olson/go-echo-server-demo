package main

import (
	"github.com/labstack/echo"
)

// ------------ root routes --------------

// definition
type rootRoutes struct {
	route
}

// mappings
func (r rootRoutes) init() {
	// avoid double forward slash in url
	// (maybe ignore this in route proxy methods)
	r.get("", r.getRoot)
}

// handlers
func (r rootRoutes) getRoot(c echo.Context) error {
	return c.String(200, "this is the root route")
}
