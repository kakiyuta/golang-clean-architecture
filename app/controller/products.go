package controller

import (
	"net/http"

	api "github.com/kakiyuta/golang-clean-architecture/app/gen"
	"github.com/labstack/echo/v4"
)

func (c *Controller) GetV1Products(ctx echo.Context, params api.GetV1ProductsParams) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Hello, World!",
	})
}
