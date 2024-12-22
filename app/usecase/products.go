package usecase

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/model"
	"github.com/kakiyuta/golang-clean-architecture/app/domain/repository"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/input"
	"github.com/kakiyuta/golang-clean-architecture/app/usecase/output"
)

type ProductsUsecase struct {
	ConnectionController db.ConnectionController
	ProductRepository    repository.Products
	VariantRepository    repository.Variants
}

func NewProductsUsecase(p repository.Products, v repository.Variants, cc db.ConnectionController) ProductsUsecase {
	return ProductsUsecase{
		ConnectionController: cc,
		ProductRepository:    p,
		VariantRepository:    v,
	}
}

// GetProducts 商品一覧を取得
func (p *ProductsUsecase) GetProducts(input input.ProductsGetProducts) (*output.ProductsGetProducts, error) {
	productsVariants, err := p.ProductRepository.GetProductsWithVariation(input.Limit, input.Offset)
	if err != nil {
		return nil, err
	}

	// トランザクションの動作確認
	p.ConnectionController.Begin()
	defer p.ConnectionController.Rollback()

	output := &output.ProductsGetProducts{
		Total:    0,
		Products: productsVariants,
	}

	p.ConnectionController.Commit()

	return output, nil
}

func (p *ProductsUsecase) CreateProduct(input input.ProductsCreateProduct) (*output.ProdunctsGreateProdunct, error) {

	output := &output.ProdunctsGreateProdunct{
		Product: model.Product{
			ID:   1,
			Name: input.Name,
		},
	}
	return output, nil
}
