package main

import (
	"fmt"

	"github.com/kakiyuta/golang-clean-architecture/app/domain/model"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
)

func main() {
	// DB接続
	db, err := db.NewMySQLConnector()
	if err != nil {
		panic(err)
	}

	product := model.Product{}
	err = db.GetSlave().First(&product, 1).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}
