package registry

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db/mysql"
)

type devRepositoryImp struct{}

func NewDevRepository() RepositoryInterface {
	return &devRepositoryImp{}
}

func (r *devRepositoryImp) NewProducts() repository.Products {
	return &mysql.Product{}
}

func (r *devRepositoryImp) NewVariants() repository.Variants {
	return &mysql.Variants{
		Con: nil,
	}
}
