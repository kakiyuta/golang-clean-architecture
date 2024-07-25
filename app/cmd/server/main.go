package main

import (
	"github.com/kakiyuta/golang-clean-architecture/app/controller"
	api "github.com/kakiyuta/golang-clean-architecture/app/gen"
	"github.com/kakiyuta/golang-clean-architecture/app/registry"
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
	con := controller.NewController(repo)
	api.RegisterHandlers(e, con)
	e.Logger.Fatal(e.Start(":1323"))

	// c := dig.New()
	// c.Provide(registry.NewLocalRepository)
	// c.Provide(controller.NewController)
	// c.Provide(registry.RepositoryInterface.NewProducts)
	// err = c.Invoke(func(con api.ServerInterface) {
	// 	fmt.Println("test")
	// 	api.RegisterHandlers(e, con)
	// 	e.Logger.Fatal(e.Start(":1323"))
	// })

	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
}
