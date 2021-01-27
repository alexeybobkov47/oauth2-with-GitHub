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
		ClientID:     "7735674",
		ClientSecret: "9qp6zpB9szRsb3ICSoat",
		Scopes:       []string{"photos"},
		RedirectURL:  "http://localhost:8080/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://oauth.vk.com/authorize",
			TokenURL: "https://oauth.vk.com/access_token",
		},
	}
	s.registerRoutes()

	// Start server

	s.l.Fatal(s.mux.Start(":8080"))

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
	e.GET("/login", func(c echo.Context) error {
		return c.Redirect(301, s.oauth.AuthCodeURL("state", oauth2.AccessTypeOffline))
	})
	e.GET("/callback", s.authorization)

}
