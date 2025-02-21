package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/kakiyuta/golang-clean-architecture/app/gen/api"
	"github.com/kakiyuta/golang-clean-architecture/app/handler/controller"
	"github.com/kakiyuta/golang-clean-architecture/app/registry"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	Init()

	e := echo.New()

	// Logger
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// APIリクエストのバリデーション
	swagger, err := api.GetSwagger()
	if err != nil {
		e.Logger.Fatal(err)
	}
	swagger.Servers = nil
	// 認証の設定
	options := &echomiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {

				switch ai.SecuritySchemeName {
				case "bearerAuth":
					zap.S().Infof("bearer token: %s", ai.RequestValidationInput.Request.Header.Get("Authorization"))

					authorization := ai.RequestValidationInput.Request.Header.Get("Authorization")

					authorizations := strings.Split(authorization, " ")

					if len(authorizations) != 2 {
						return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("invalid authorization: %s", authorization))
					}

					if authorizations[0] != "Bearer" {
						return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("invalid token: %s", authorization))
					}

					if authorizations[1] != "token" {
						return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("invalid token: %s", authorizations[1]))
					}
				default:
					return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("unknown security scheme: %s", ai.SecuritySchemeName))
				}

				return nil
			},
		},
	}

	e.Use(echomiddleware.OapiRequestValidatorWithOptions(swagger, options))

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
	defer func() {
		_ = logger.Sync() //nolint:errcheck
	}()
	zap.ReplaceGlobals(logger)
}
