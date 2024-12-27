package registry

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db/moc"
)

type localRepositoryImp struct {
	db *db.LocalConnector
}

func NewLocalRepository() RepositoryInterface {
	db := db.NewLocalConnector()
	return &localRepositoryImp{db}
}

func (r *localRepositoryImp) GetDB() db.ConnectionController {
	return r.db
}

func (r *localRepositoryImp) NewProducts() repository.Products {
	return &moc.Product{}
}

func (r *localRepositoryImp) NewVariants() repository.Variants {
	return &moc.Variants{}
}
