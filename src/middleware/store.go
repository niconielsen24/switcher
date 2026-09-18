package middleware

import (
	"switcher/store"

	"github.com/labstack/echo/v5"
)

func NewGameStore(s store.GameStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("gameStore", s)
			return next(c)
		}
	}
}

func NewUserStore(s store.UserStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("userStore", s)
			return next(c)
		}
	}
}
