package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/kakiyuta/golang-clean-architecture/app/gen/grpc"
	"github.com/kakiyuta/golang-clean-architecture/app/handler/rpc"
	"github.com/kakiyuta/golang-clean-architecture/app/registry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	Init()

	// 1. 8080番portのListenerを作成
	port := 1323
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	// s := grpc.NewServer()
	s := ggrpc.NewServer()

	// 3. gRPCサーバーにGreetingServiceを登録
	repo := registry.NewDevRepository()
	grpc.RegisterGreetingServiceServer(s, rpc.NewProductServer(repo))

	reflection.Register(s)

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		err := s.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()

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
