package api

import (
	"switcher/types"

	"github.com/labstack/echo/v5"
)

func NewUser(c *echo.Context) error {
	return c.JSON(200, types.NewUser("user"))
}
