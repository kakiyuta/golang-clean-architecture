package main

import (
	"github.com/kakiyuta/golang-clean-architecture/app/controller"
	"github.com/kakiyuta/golang-clean-architecture/app/gen/api"
	"github.com/kakiyuta/golang-clean-architecture/app/registry"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	smiddleware "github.com/oapi-codegen/echo-middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	Init()

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
	repo := registry.NewDevRepository()
	con := controller.NewController(repo)
	api.RegisterHandlers(e, con)
	e.Logger.Fatal(e.Start(":1323"))
}

func Init() {
	// ロガーの初期化
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	conf := zap.Config{
		Level:    level,
		Encoding: "console", // json or console
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "name",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := conf.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}
