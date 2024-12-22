package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQLConnector struct {
	Connection    *gorm.DB
	tranzaction   *gorm.DB
	isTranzaction bool
}

func NewMySQLConnector() (*MySQLConnector, error) {
	// TODO : 環境変数から取得するように変更すること
	sqlDB, err := sql.Open("mysql", "user:password@tcp(ec-db:3306)/ec?charset=utf8&parseTime=true")

	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logSQLError(),
	})
	if err != nil {
		return nil, err
	}

	return &MySQLConnector{
		Connection: db,
	}, nil
}

// GetMaster Master DB接続を取得する
func (c *MySQLConnector) GetMaster() *gorm.DB {
	return c.Connection
}

// GetSlave Slave DB接続を取得する
func (c *MySQLConnector) GetSlave() *gorm.DB {
	return c.Connection
}

func (c *MySQLConnector) Begin() error {
	if c.isTranzaction {
		return errors.New("トランザクションが既に開始されています")
	}
	c.tranzaction = c.Connection.Begin()
	return nil
}

func (c *MySQLConnector) Commit() error {
	if !c.isTranzaction {
		return errors.New("トランザクションが開始されていません")
	}
	c.tranzaction.Commit()

	c.tranzaction = nil
	return nil
}

func (c *MySQLConnector) Rollback() {
	if !c.isTranzaction {
		return
	}
	c.tranzaction = c.tranzaction.Rollback()
	if c.tranzaction.Error != nil {
		log.Println("ロールバック中にエラー発生", c.tranzaction.Error)
	}
	c.tranzaction = nil
}

func logSQLError() logger.Interface {
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
