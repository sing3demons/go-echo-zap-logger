package healthchk

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sing3demons/echo-logger/mlog"
)

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (h handler) Check(c echo.Context) error {
	logger := mlog.L(c)
	logger.Info("health check")

	return c.String(http.StatusOK, "hey Gopher!, I'm alive!")
}
