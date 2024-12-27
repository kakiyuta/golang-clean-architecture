package main

import (
	"github.com/kakiyuta/golang-clean-architecture/app/controller"
	"github.com/kakiyuta/golang-clean-architecture/app/gen/api"
	"github.com/kakiyuta/golang-clean-architecture/app/registry"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	smiddleware "github.com/oapi-codegen/echo-middleware"
)

func main() {
	// // 環境変数の読み込み
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

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
	// repo := registry.NewLocalRepository()
	repo := registry.NewDevRepository()
	con := controller.NewController(repo)
	api.RegisterHandlers(e, con)
	e.Logger.Fatal(e.Start(":1323"))

}
