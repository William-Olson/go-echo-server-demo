package main

import (
	"github.com/labstack/echo"
)

// ----------- routing helper setup ---------------------

type someRoutes interface {
	init()
}

func mountRoutes(r someRoutes) {
	r.init()
}

// the Parent/Promoted struct of routes is a router

type route struct {
	base string
	echo *echo.Echo
}

// proxy methods for terseness
func (r route) get(sub string, cb func(echo.Context) error) {
	r.echo.GET(r.base+sub, cb)
}
func (r route) post(sub string, cb func(echo.Context) error) {
	r.echo.POST(r.base+sub, cb)
}
func (r route) put(sub string, cb func(echo.Context) error) {
	r.echo.PUT(r.base+sub, cb)
}
func (r route) delete(sub string, cb func(echo.Context) error) {
	r.echo.DELETE(r.base+sub, cb)
}
