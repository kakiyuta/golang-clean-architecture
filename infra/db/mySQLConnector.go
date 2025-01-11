package db

import (
	"database/sql"
	"errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnector struct {
	connection    *gorm.DB
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
		// Logger: logSQLError(),
	})
	if err != nil {
		return nil, err
	}

	return &MySQLConnector{
		connection: db,
	}, nil
}

// GetMaster Master DB接続を取得する
func (c *MySQLConnector) GetMaster() *gorm.DB {
	if c.isTranzaction {
		return c.tranzaction
	}
	return c.connection
}

// GetSlave Slave DB接続を取得する
func (c *MySQLConnector) GetSlave() *gorm.DB {
	if c.isTranzaction {
		return c.tranzaction
	}
	return c.connection
}

func (c *MySQLConnector) Begin() error {
	log.Printf("Begin-----------------------")
	if c.isTranzaction {
		return errors.New("トランザクションが既に開始されています")
	}
	c.tranzaction = c.connection.Begin()
	c.isTranzaction = true
	return nil
}

func (c *MySQLConnector) Commit() error {
	if !c.isTranzaction {
		return errors.New("トランザクションが開始されていません")
	}
	log.Printf("Commit-----------------------")
	c.tranzaction.Commit()

	c.isTranzaction = false
	c.tranzaction = nil
	return nil
}

func (c *MySQLConnector) Rollback() {
	if !c.isTranzaction {
		return
	}
	log.Printf("Rollback-----------------------")
	c.tranzaction = c.tranzaction.Rollback()
	if c.tranzaction.Error != nil {
		log.Println("ロールバック中にエラー発生", c.tranzaction.Error)
	}

	c.isTranzaction = false
	c.tranzaction = nil
}
