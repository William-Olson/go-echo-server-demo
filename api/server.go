package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"models"
)

const (
	jwtSecret  = "verygoodsecret"
	serverPort = "7447"
)

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

type pmap struct {
	payload map[string](interface{})
}

/*

	Init routes and start the server

*/
func (s *Server) Start() {

	s.e = echo.New()

	// some general middleware
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	// create the route groups
	rootGroup := s.e.Group("")
	userGroup := s.e.Group("/users")

	// auth middleware
	userGroup.Use(middleware.JWT([]byte(jwtSecret)))

	// define base paths
	routes := []someRoutes{
		rootRoutes{route{rootGroup, s.Db}},
		userRoutes{route{userGroup, s.Db}},
	}

	// wire up sub paths
	for _, r := range routes {
		r.mount()
	}

	s.e.Logger.Fatal(s.e.Start(":" + serverPort))

}

/*

	Response payload helpers

*/

func newResponse(s string, v interface{}) pmap {

	payload := pmap{(map[string](interface{}){})}
	payload.payload[s] = v
	return payload

}

func (p *pmap) set(s string, v interface{}) {

	p.payload[s] = v

}
