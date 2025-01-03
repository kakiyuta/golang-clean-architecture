package controller

import (
	"net/http"

	"github.com/kakiyuta/golang-clean-architecture/app/gen/api"
	"github.com/kakiyuta/golang-clean-architecture/app/library/weberrors"
	"github.com/kakiyuta/golang-clean-architecture/app/registry"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	repo registry.RepositoryInterface
}

func NewController(repo registry.RepositoryInterface) api.ServerInterface {
	return &Controller{
		repo: repo,
	}
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func StringPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

// errorResponse エラーレスポンスを返す
func errorResponse(ctx echo.Context, err error) error {
	httpCode := http.StatusInternalServerError
	if we, ok := err.(*weberrors.WebError); ok {
		httpCode = we.Code
	}

	return ctx.JSON(httpCode, map[string]interface{}{
		"code": httpCode,
		"msg":  err.Error(),
	})
}
