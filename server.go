package main

import (
	"github.com/labstack/echo"
)

type Server struct {
	db *DB
	e  *echo.Echo
}

func (s *Server) start() {
	s.e = echo.New()

	// route setup
	mountRoutes(rootRoutes{route{"/", s.e, s.db}})
	mountRoutes(userRoutes{route{"/users", s.e, s.db}})

	s.e.Logger.Fatal(s.e.Start(":7447"))
}
