package api

import (
	"github.com/labstack/echo"
	"models"
)

type Server struct {
	Db *models.DB
	e  *echo.Echo
}

func (s *Server) Start() {
	s.e = echo.New()

	// route setup
	mountRoutes(rootRoutes{route{"/", s.e, s.Db}})
	mountRoutes(userRoutes{route{"/users", s.e, s.Db}})

	s.e.Logger.Fatal(s.e.Start(":7447"))
}
