package usecase

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/input"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/output"
)

type ProductsUsecase struct {
	ProductRepository repository.Products
}

func NewProductsUsecase(p repository.Products) ProductsUsecase {
	return ProductsUsecase{
		ProductRepository: p,
	}
}

func (p *ProductsUsecase) GetProducts(input input.ProductsGetProducts) (*output.ProductsGetProducts, error) {
	productsVariants, err := p.ProductRepository.GetProductsWithVariation(input.Limit, input.Offset)
	if err != nil {
		return nil, err
	}
	output := &output.ProductsGetProducts{
		Total:    0,
		Products: productsVariants,
	}
	return output, nil
}
