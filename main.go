package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
)

type server struct {
	mux   *echo.Echo
	l     echo.Logger
	oauth *oauth2.Config
}

func main() {
	s := &server{}
	s.oauth = &oauth2.Config{
		ClientID:     "7fd1e335e67efe6e8674",
		ClientSecret: "17vPWEL%",
		Scopes:       []string{"test"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	s.registerRoutes()

	// Start server

	s.l.Fatal(s.mux.Start(":1323"))

}

func (s *server) registerRoutes() {
	// Echo instance
	e := echo.New()
	s.mux = e
	s.l = e.Logger

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.Static("/", "web")
	// e.GET("/login", func(c echo.Context) error {
	// 	return c.Redirect(301, s.oauth.AuthCodeURL("state", oauth2.AccessTypeOnline))
	// })
	// e.GET("/authorization", s.authorization)
	e.GET("/login", s.authorization)

}
