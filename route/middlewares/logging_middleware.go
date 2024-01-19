package middlewares

import (
	"echo-project/logger"
	"github.com/labstack/echo/v4"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		l := logger.Get()
		l.Info().
			Str("method", c.Request().Method).
			Str("uri", c.Request().URL.Path).
			Str("query", c.Request().URL.RawQuery).
			Msg("Request")

		// call the next middleware/handler
		err := next(c)
		if err != nil {
			l.Error().
				Str("error", err.Error()).
				Msg("Response")
			return err
		}
		return nil
	}
}
