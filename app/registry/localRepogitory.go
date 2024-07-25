package registry

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db/moc"
)

type localRepositoryImp struct{}

func NewLocalRepository() RepositoryInterface {
	return &localRepositoryImp{}
}

func (r *localRepositoryImp) NewProducts() repository.Products {
	return &moc.Product{}
}

func (r *localRepositoryImp) NewVariants() repository.Variants {
	return &moc.Variants{}
}
