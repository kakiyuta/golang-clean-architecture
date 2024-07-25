package usecase

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/input"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/output"
)

type ProductsUsecase struct {
	ProductRepository repository.Products
	VariantRepository repository.Variants
}

func NewProductsUsecase(p repository.Products, v repository.Variants) ProductsUsecase {
	return ProductsUsecase{
		ProductRepository: p,
		VariantRepository: v,
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
