package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func ParseHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header != "" {
			key := strings.Split(header, "Bearer")[1]
			trimmedKey := strings.TrimSpace(key)
			c.Set("key", trimmedKey)
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
		return next(c)
	}
}
