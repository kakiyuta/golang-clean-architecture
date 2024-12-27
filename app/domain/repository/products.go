package repository

import "github.com/kakiyuta/golang-clean-architecture/app/domain/model"

type Products interface {
	GetProductsWithVariation(limit int, offset int) ([]model.Product, error)
	GetProducts(limit int, offset int) ([]model.Product, error)
	GetProductByID(id int) (model.Product, error)
	CreateProduct(produnct model.Product) (model.Product, error)
}
