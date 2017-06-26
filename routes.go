package main

import (
	"github.com/labstack/echo"
	"regexp"
)

type someRoutes interface {
	init()
}

// route has a base path, and echo & db reference
type route struct {
	base string
	echo *echo.Echo
	db   *DB
}

// call the init method of someRoutes the init method should wire up the
// someRoutes paths via get post put delete proxy methods
func mountRoutes(r someRoutes) {

	r.init()

}

// proxy methods

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

// helper for joining url base and tail paths
func slashJoin(s1 string, s2 string) string {

	rg := regexp.MustCompile("//*")
	return rg.ReplaceAllString(s1+s2, "/")

}
