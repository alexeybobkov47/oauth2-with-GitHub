package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func (s *server) authorization(c echo.Context) error {
	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	state := c.FormValue("state")
	if state != "state" {
		return echo.ErrBadRequest
	}
	code := c.FormValue("code")
	tok, err := s.oauth.Exchange(oauth2.NoContext, code)
	fmt.Printf("\n%v", tok)
	if err != nil {
		return err
	}
	getInfoURL := "https://api.vk.com/method/friends.get?fields=?&access_token=" + tok.AccessToken + "&v=5.126"
	c.Redirect(301, getInfoURL)

	return err

}
