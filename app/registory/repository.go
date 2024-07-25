package registry

import "github.com/kakiyuta/golang-clean-architecture/app/domain/repository"

type RepositoryInterface interface {
	NewProducts() repository.Products
}
