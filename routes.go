package main

import (
	"github.com/labstack/echo"
	"regexp"
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
	db   *DB
}

func slashJoin(s1 string, s2 string) string {
	rg := regexp.MustCompile("//*")
	return rg.ReplaceAllString(s1+s2, "/")
}

// proxy methods for terseness
func (r route) get(sub string, cb func(echo.Context) error) {
	r.echo.GET(slashJoin(r.base, sub), cb)
}
func (r route) post(sub string, cb func(echo.Context) error) {
	r.echo.POST(slashJoin(r.base, sub), cb)
}
func (r route) put(sub string, cb func(echo.Context) error) {
	r.echo.PUT(slashJoin(r.base, sub), cb)
}
func (r route) delete(sub string, cb func(echo.Context) error) {
	r.echo.DELETE(slashJoin(r.base, sub), cb)
}
