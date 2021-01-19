package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func (s *server) authorization(c echo.Context) error {
	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	code := c.FormValue("code")
	fmt.Printf("\n%v\n", code)
	tok, err := s.oauth.Exchange(c.Request().Context(), code)
	fmt.Printf("\n%v\n", tok)
	if err != nil {
		log.Fatal(err)
	}

	// client := s.oauth.Client(c.Request().Context(), tok)
	// client.Get("...")

	return err

}
