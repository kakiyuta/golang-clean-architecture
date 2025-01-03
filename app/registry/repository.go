package registry

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/dao/mysql"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
)

type RepositoryInterface interface {
	GetDB() db.Connector
	NewProducts() repository.Products
	NewVariants() repository.Variants
}

type repositoryImp struct {
	db *db.MySQLConnector
}

func NewDevRepository() RepositoryInterface {
	db, err := db.NewMySQLConnector()
	if err != nil {
		panic(err)
	}
	return &repositoryImp{db: db}
}

func (r *repositoryImp) GetDB() db.Connector {
	return r.db
}

func (r *repositoryImp) NewProducts() repository.Products {
	return &mysql.Product{
		Con: r.db,
	}
}

func (r *repositoryImp) NewVariants() repository.Variants {
	return &mysql.Variants{
		Con: r.db,
	}
}
