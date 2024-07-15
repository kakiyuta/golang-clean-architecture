package main

import (
	"github.com/kakiyuta/golang-clean-architecture/app/controller"
	api "github.com/kakiyuta/golang-clean-architecture/app/gen"
	"github.com/labstack/echo/v4"
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

	controller := controller.Controller{}
	api.RegisterHandlers(e, &controller)

	e.Logger.Fatal(e.Start(":1323"))
}
