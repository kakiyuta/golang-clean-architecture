package mysql

import (
	"context"
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnector struct {
	// connectionContext    rdbms.ConnectionContext
	isOpen bool
	// masterConnectionType masterConnectionType
	dbSlave       *gorm.DB
	dbMaster      *gorm.DB
	dbTransaction *gorm.DB
	hasContext    bool
	ctx           context.Context
}

// NewMySQLConnector MySQLConnectorを生成する
func NewMySQLConnector() *MySQLConnector {
	c := MySQLConnector{}
	c.setMaster()
	c.setSlave()
	return &c
}

// GetSlave SlaveDBへの接続を取得する
func (c *MySQLConnector) setSlave() {
	sqlDB, err := sql.Open("mysql", "user:password@tcp(ec-db:3306)/ec?charset=utf8&parseTime=true")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))

	if err != nil {
		panic(err)
	}

	c.dbSlave = db
}

// GetMaster MasterDBへの接続を取得する
func (c *MySQLConnector) setMaster() {
	sqlDB, err := sql.Open("mysql", "user:password@tcp(ec-db:3306)/ec?charset=utf8&parseTime=true")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))

	if err != nil {
		panic(err)
	}

	c.dbMaster = db
}

// Open 接続を開く
func (c *MySQLConnector) Open() error {
	c.isOpen = true
	return nil
}

// Close 接続を閉じる
func (c *MySQLConnector) Close() {
	c.isOpen = false
}

// Begin トランザクションを開始する
func (c *MySQLConnector) Begin() error {
	// TODO: トランザクションの開始処理を実装する
	return nil
}

// Commit トランザクションをコミットし、接続を閉じる
func (c *MySQLConnector) Commit() error {
	// TODO: トランザクションのコミット処理を実装する
	return nil
}

// Rollback トランザクションをロールバックし、接続を閉じる
func (c *MySQLConnector) Rollback() {
	// TODO: トランザクションのロールバック処理を実装する
}
