package registry

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db/mysql"
)

type devRepositoryImp struct {
	db *db.MySQLConnector
}

func NewDevRepository() RepositoryInterface {
	db, err := db.NewMySQLConnector()
	if err != nil {
		panic(err)
	}
	return &devRepositoryImp{db: db}
}

func (r *devRepositoryImp) NewProducts() repository.Products {
	return &mysql.Product{
		Con: r.db.Connection,
	}
}

func (r *devRepositoryImp) NewVariants() repository.Variants {
	return &mysql.Variants{
		Con: r.db.Connection,
	}
}
