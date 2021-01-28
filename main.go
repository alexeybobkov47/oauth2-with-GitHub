package main

import (
	"github.com/casbin/casbin/v2"
	casbin_mw "github.com/labstack/echo-contrib/casbin"
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
		Scopes:       []string{},
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
	enforcer, err := casbin.NewEnforcer("casbin_auth_model.conf", "casbin_auth_policy.csv")
	if err != nil {
		s.l.Fatalf("error enforcer: %v", err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(casbin_mw.Middleware(enforcer))

	// Routes
	e.Static("/", "web")
	e.GET("/login", func(c echo.Context) error {
		return c.Redirect(301, s.oauth.AuthCodeURL("state", oauth2.AccessTypeOffline))
	})
	e.GET("/callback", s.authorization)
	e.File("/admin", "web/admin.html")
}
