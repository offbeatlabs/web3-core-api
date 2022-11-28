package middleware

import (
	"github.com/arhamj/go-commons/pkg/logger"
	"github.com/labstack/echo/v4"
	"time"
)

func LoggingMiddleware(l *logger.AppLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			l.HttpMiddlewareAccessLogger(req.Method, req.RequestURI, res.Status, res.Size, time.Since(start))
			return nil
		}
	}
}
