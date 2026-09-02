package api

import (
	"github.com/labstack/echo/v4"
)

func HandleGetGame(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"message": "Game retrieved successfully",
	})
}
