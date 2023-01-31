package middleware

import (
	"github.com/arhamj/go-commons/pkg/constants"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"time"
)

func LoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			HttpMiddlewareAccessLogger(req.Method, req.RequestURI, res.Status, res.Size, time.Since(start))
			return nil
		}
	}
}

func HttpMiddlewareAccessLogger(method, uri string, status int, size int64, time time.Duration) {
	log.WithFields(log.Fields{
		constants.METHOD: method,
		constants.URI:    uri,
		constants.STATUS: status,
		constants.SIZE:   size,
		constants.TIME:   time,
	}).Info(
		constants.HTTP,
	)
}
