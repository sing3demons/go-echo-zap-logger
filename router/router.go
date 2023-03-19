package router

import (
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"

	"github.com/sing3demons/echo-logger/healthchk"
	"github.com/sing3demons/echo-logger/mlog"
	"go.uber.org/zap"
)

func RegRoute(logger *zap.Logger) *echo.Echo {
	e := echo.New()

	e.Use(mlog.Middleware(logger))
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	hHealthChk := healthchk.New()
	e.GET("/healthz", hHealthChk.Check)

	e.GET("/", func(c echo.Context) error {
		return response(c, http.StatusOK, "Hello, World!")
	})

	e.GET("/ping", func(c echo.Context) error {
		msg := map[string]interface{}{"ping": "pong"}
		return response(c, http.StatusOK, msg)
	})

	return e
}

func response(c echo.Context, code int, value interface{}) error {
	logger := mlog.L(c)
	logger.Info("response", zap.Any("response", value))
	return c.JSON(code, value)
}
