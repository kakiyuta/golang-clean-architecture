package registry

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
)

type RepositoryInterface interface {
	GetDB() db.Connector
	NewProducts() repository.Products
	NewVariants() repository.Variants
}
