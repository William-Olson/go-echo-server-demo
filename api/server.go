package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"models"
)

const jwtSecret = "verygoodsecret"

type Server struct {
	Db *models.DB
	e  *echo.Echo
}

type someRoutes interface {
	mount()
}

type route struct {
	group *echo.Group
	db    *models.DB
}

/*

	Init routes and start the server

*/
func (s *Server) Start() {

	s.e = echo.New()

	// some general middleware
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	// define base paths
	routes := []someRoutes{
		rootRoutes{route{s.e.Group(""), s.Db}},
		userRoutes{route{s.e.Group("/users"), s.Db}},
	}

	// wire up sub paths
	for _, r := range routes {
		r.mount()
	}

	s.e.Logger.Fatal(s.e.Start(":7447"))

}
