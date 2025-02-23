package controller

import (
	"net/http"

	"github.com/kakiyuta/golang-clean-architecture/app/domain/dto/input"
	"github.com/kakiyuta/golang-clean-architecture/app/gen/api"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase"
	"github.com/labstack/echo/v4"
)

type loginResult struct {
	Token string `json:"token"`
}

func (c *Controller) PostV1Login(ctx echo.Context) error {
	var accessUser api.PostV1LoginJSONRequestBody
	if err := ctx.Bind(&accessUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	input := input.NewLoginRequestUser(string(accessUser.Email), accessUser.Password)
	usecase := c.newAuthUseCase()
	oputput, err := usecase.Login(input)
	if err != nil {
		return errorResponse(ctx, err)
	}

	// ログイン成功時の処理
	return ctx.JSON(http.StatusOK, loginResult{
		Token: oputput.Token,
	})
}

func (c *Controller) newAuthUseCase() usecase.AuthUsecase {
	return usecase.NewAuthUsecase(c.repo.GetDB())
}
