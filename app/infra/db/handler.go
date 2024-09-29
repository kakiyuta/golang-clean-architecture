package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() (*SqlHandler, error) {
	// TODO : 環境変数から取得するように変更すること
	sqlDB, err := sql.Open("mysql", "user:password@tcp(ec-db:3306)/ec?charset=utf8&parseTime=true")

	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: LogSQLError(),
	})
	if err != nil {
		return nil, err
	}

	return &SqlHandler{
		Conn: db,
	}, nil
}

func LogSQLError() logger.Interface {
	logLevel := logger.Info
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
}
