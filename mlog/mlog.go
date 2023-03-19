package mlog

import (
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const key = "logger"

func L(c echo.Context) *zap.Logger {
	switch logger := c.Get(key).(type) {
	case *zap.Logger:
		return logger
	default:
		return zap.NewNop()
	}
}

func Middleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return logMiddleware(next, logger)
	}
}

func logMiddleware(next echo.HandlerFunc, logger *zap.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		l := setLog(c, logger)
		c.Set(key, l)
		return next(c)
	}
}

func setLog(c echo.Context, logger *zap.Logger) *zap.Logger {
	xParent := c.Request().Header.Get("X-Parent-ID")
	userAgent := c.Request().UserAgent()
	hostname, _ := os.Hostname()
	start := time.Now()
	path := c.Request().URL.Path
	query := c.Request().URL.RawQuery
	status := strconv.Itoa(c.Response().Status)

	if xParent == "" {
		xParent = uuid.NewString()
	}
	xSpan := uuid.NewString()
	return logger.With(
		zap.String("parent-id", xParent),
		zap.String("span-id", xSpan),
		zap.String("User-Agent", userAgent),
		zap.String("hostname", hostname),
		zap.String("ip", c.RealIP()),
		zap.String("status", start.String()),
		zap.String("path", path),
		zap.String("query", query),
		zap.String("status", status),
		zap.String("method", c.Request().Method),
	)
}
