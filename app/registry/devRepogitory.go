package registry

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db/mysql"
)

type devRepositoryImp struct {
	db *db.SqlHandler
}

func NewDevRepository() RepositoryInterface {
	db, err := db.NewSqlHandler()
	if err != nil {
		panic(err)
	}
	return &devRepositoryImp{db: db}
}

func (r *devRepositoryImp) NewProducts() repository.Products {
	return &mysql.Product{}
}

func (r *devRepositoryImp) NewVariants() repository.Variants {
	return &mysql.Variants{
		Con: nil,
	}
}
