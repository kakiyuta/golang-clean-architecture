package registry

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db/moc"
)

type LocalRepositoryImp struct{}

func NewLocalRepository() *LocalRepositoryImp {
	return &LocalRepositoryImp{}
}

func (r *LocalRepositoryImp) NewProducts() repository.Products {
	return moc.Product{}
}
