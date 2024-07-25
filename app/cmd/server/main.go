package main

import (
	"github.com/kakiyuta/golang-clean-architecture/app/controller"
	api "github.com/kakiyuta/golang-clean-architecture/app/gen"
	registry "github.com/kakiyuta/golang-clean-architecture/app/registory"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	smiddleware "github.com/oapi-codegen/echo-middleware"
)

func main() {
	e := echo.New()

	// APIリクエストのバリデーション
	swagger, err := api.GetSwagger()
	if err != nil {
		e.Logger.Fatal(err)
	}
	swagger.Servers = nil
	e.Use(smiddleware.OapiRequestValidator(swagger))

	// Logger
	e.Use(middleware.Logger())

	// ローカルテスト用のレポジトリを作成
	repo := registry.NewLocalRepository()
	controller := controller.NewController(repo)
	api.RegisterHandlers(e, &controller)

	e.Logger.Fatal(e.Start(":1323"))
}
